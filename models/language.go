package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	LanguageName string
	LanguageCode string
	Translations []Translation
	TextContent []TextContent
}
