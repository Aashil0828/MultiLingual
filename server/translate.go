package server

import (
	"context"
	"errors"
	"fmt"
	"multilingual-new/models"
	"multilingual-new/pb/pb"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type MultiLingualServer struct {
	Db     *gorm.DB
	Client *translate.Client
	pb.UnimplementedMultiLingualServiceServer
}

func (m *MultiLingualServer) Translate(ctx context.Context, req *pb.MultiLingualRequest) (*pb.MultiLingualResponse, error) {
	tx := m.Db.Begin()
	var Data []*pb.TextContent
	for _, TextContent := range req.Data {
		var Translation string
		var lang_code string
		var pbTextContent *pb.TextContent
		if tx.Model(&models.TextContent{}).Select("translation").Joins("join translations on translations.text_content_id = text_contents.id").Where("text_contents.original_text = ? AND translations.language_id = ?", TextContent.Label, req.GetLanguageId()).First(&Translation).RowsAffected != 0 {
			Data = append(Data, &pb.TextContent{Id: TextContent.GetId(), Label: Translation})
		} else {
			fmt.Print("hi")
			if tx.Model(&models.Language{}).Select("language_code").Find(&lang_code, req.GetLanguageId()).RowsAffected == 0 {
				tx.Rollback()
				return &pb.MultiLingualResponse{Status: int32(codes.NotFound), Message: fmt.Sprintf("could not find language code for language id %v", req.GetLanguageId())}, errors.New("could not find language code for language id")
			} else {
				lang, err := language.Parse(lang_code)
				if err != nil {
					tx.Rollback()
					return &pb.MultiLingualResponse{Status: int32(codes.Internal), Message: fmt.Sprintf("could not parse language code : %v", err)}, fmt.Errorf("language cannot be Parsed: %v", err)
				}
				resp, err := m.Client.Translate(ctx, []string{TextContent.Label}, lang, &translate.Options{
					Model: "nmt",
				})
				if err != nil {
					return &pb.MultiLingualResponse{Status: int32(codes.Internal), Message: fmt.Sprintf("could not translate content : %v", err)}, fmt.Errorf("cannot translate text: %v", err)
				}
				if len(resp) == 0 {
					return &pb.MultiLingualResponse{Status: int32(codes.Internal), Message: fmt.Sprintf("could not translate content : %v", err)}, nil
				}
				var orig_lang_id uint
				fmt.Println(resp[0].Source.String())
				if tx.Model(&models.Language{}).Select("id").Where(&models.Language{LanguageCode: resp[0].Source.String()}).Find(&orig_lang_id).RowsAffected == 0 {
					tx.Rollback()
					return &pb.MultiLingualResponse{Status: int32(codes.NotFound), Message: "input text not in supported languages"}, errors.New("input text not in supported languages")
				} else {
					if tx.Create(&models.TextContent{OriginalText: TextContent.Label, LanguageId: orig_lang_id, Translations: []models.Translation{{LanguageId: uint(req.GetLanguageId()), Translation: resp[0].Text}}}).RowsAffected == 0 {
						tx.Rollback()
						return &pb.MultiLingualResponse{Status: int32(codes.Internal), Message: "cannot insert text into db"}, errors.New("cannot insert text into db")
					}
					fmt.Println(resp)
					pbTextContent = &pb.TextContent{Label: resp[0].Text, Id: TextContent.GetId()}
				}
			}
			Data = append(Data, pbTextContent)
		}
	}
	tx.Commit()
	return &pb.MultiLingualResponse{Status: int32(codes.OK), Message: "translated successfully", Data: Data}, nil
}
