package gorms

import (
	"unisun/api/authen-listening/src/config"
	"unisun/api/authen-listening/src/entitys"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAuthPermission interface {
	FindAndCreate(Data entitys.UserAuthPermission) *gorm.DB
}

type dbServices struct {
	Context *gorm.DB
}

func JWTAuthService() UserAuthPermission {
	return &dbServices{
		Context: config.DB,
	}
}

func (db *dbServices) FindAndCreate(Data entitys.UserAuthPermission) *gorm.DB {
	return db.Context.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Data)
}
