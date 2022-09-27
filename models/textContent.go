package models

import "gorm.io/gorm"

type TextContent struct {
	gorm.Model
	OriginalText string
	LanguageId uint
	Translations []Translation
}

