package models

import "gorm.io/gorm"

type Translation struct {
	gorm.Model
	TextContentId uint
	LanguageId uint
	Translation string
}