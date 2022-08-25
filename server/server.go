package server

import (
	"context"
	"database/sql"
	"fmt"
	"multilingual-new/pb/pb"
	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type Server struct {
	Db *sql.DB
	pb.UnimplementedMultiLingualServiceServer
}

func (s *Server) Translate(ctx context.Context, req *pb.MultiLingualRequest) (*pb.MultiLingualResponse, error) {
	tx, err:= s.Db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		return nil, fmt.Errorf("cannot start transaction: %v", err)
	}
	result, err := s.Db.Query(`select translation from public."Translations" join public."TextContent" on public."TextContent".text_content_id=public."Translations".text_content_id join public."Languages" on public."Languages".language_id=public."Translations".language_id where lower(public."TextContent".original_text) = lower($1) and lower(public."Languages".language_name) = lower($2)`, req.Text, req.Language)
	if err != nil {
		return &pb.MultiLingualResponse{}, fmt.Errorf("cannot execute find translation query: %v", err)
	}
	var translation string
	if !result.Next() {
		ctx := context.Background()
		lang_code_query := `select language_code, language_id from public."Languages" where lower(language_name) = lower($1)`
		var lang_code string
		var target_lang_id string
		var orig_lang_id string
		var text_content_id string
		rows, err:=s.Db.Query(lang_code_query, req.GetLanguage())
		if err != nil {
			return nil, err
		}
		if rows.Next(){
			err = rows.Scan(&lang_code, &target_lang_id)
			if err != nil {
				return nil, err
			}
		} else{
			return nil, fmt.Errorf("language currently not supported: %v", req.GetLanguage())
		}
		lang, err := language.Parse(lang_code)
		if err != nil {
			return nil, fmt.Errorf("language cannot be Parsed: %v", err)
		}
		client, err := translate.NewClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("translation NewClient cannot be created: %v", err)
		}
		defer client.Close()
		resp, err := client.Translate(ctx, []string{req.GetText()}, lang, &translate.Options{
			Model: "nmt",
		})
		if err != nil {
			return nil, fmt.Errorf("cannot translate text: %v", err)
		}
		if len(resp) == 0 {
			return nil, nil
		}
		original_lang_id_query := `select language_id from public."Languages" where language_name = 'english'`
		result, err := s.Db.Query(original_lang_id_query)
		if err != nil {
			return nil, fmt.Errorf("cannot retrive original language id: %v", err)
		}
		for result.Next(){
			err = result.Scan(&orig_lang_id)
			if err != nil {
				return nil, fmt.Errorf("cannot scan orig_lang_id: %v", err)
			}
			fmt.Println(orig_lang_id)
		}
		insertTextQuery:= `insert into public."TextContent" values(default, $1, $2)`
		_, err = tx.ExecContext(ctx, insertTextQuery, req.GetText(), orig_lang_id)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("cannot insert text into TextContent : %v", err)
		}
		textContentidQuery:=`select text_content_id from public."TextContent" where lower(original_text) = lower($1)`
		textContentidResult, err:= tx.Query(textContentidQuery, req.GetText())
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve text content id : %v", err)
		}
		for textContentidResult.Next(){
			err = textContentidResult.Scan(&text_content_id)
			if err != nil {
				return nil, fmt.Errorf("cannot scan text_content_id: %v", err)
			}
			fmt.Println(text_content_id)
		}
		insertTranslationQuery:= `insert into public."Translations" values(default, $1, $2, $3)`
		_, err = tx.ExecContext(ctx, insertTranslationQuery, text_content_id, target_lang_id, resp[0].Text)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("cannot insert translation into Translations : %v", err)
		}
		translation = resp[0].Text
	} else {
		err = result.Scan(&translation)
		if err != nil {
			return nil, fmt.Errorf("cannot scan into translation: %v", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("cannot commit transaction: %v", err)
	}
	return &pb.MultiLingualResponse{Text_In_That_Language: translation}, nil
}
