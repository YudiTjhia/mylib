package model

import "github.com/jinzhu/gorm"

type BaseGormModel struct {
	DB        *gorm.DB
	TableName string
}
