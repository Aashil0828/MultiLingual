package server

import (
	"context"
	"errors"
	"multilingual-new/models"
	"multilingual-new/pb/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (m *MultiLingualServer) GetSupportedLanguages(ctx context.Context, in *emptypb.Empty) (*pb.GetSupportedLanguagesResponse, error) {
	var Languages []*models.Language
	result := m.Db.Find(&Languages)
	if result.RowsAffected == 0 {
		return &pb.GetSupportedLanguagesResponse{Status: int32(codes.NotFound), Message: "Cannot find any language"}, errors.New("languages cannot be found")
	}
	return &pb.GetSupportedLanguagesResponse{Status: int32(codes.OK), Message: "get languages support successful", Data: Loadmodelintopblang(Languages)}, nil
}

func Loadmodelintopblang(languages []*models.Language) []*pb.Language {
	var Languages []*pb.Language
	for _, language := range languages {
		Languages = append(Languages, &pb.Language{Id: uint32(language.ID), LanguageName: language.LanguageName})
	}
	return Languages
}
