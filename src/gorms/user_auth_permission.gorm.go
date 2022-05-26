package gorms

import (
	"unisun/api/authen-listening/src/config"
	"unisun/api/authen-listening/src/entitys"

	"gorm.io/gorm"
)

type UserAuthPermission interface {
	FindbyUserid(userId int) entitys.UserAuthPermission
	Create(Data entitys.UserAuthPermission)
	UpdateVersionToken(versionToken int, Data entitys.UserAuthPermission)
}

type dbServices struct {
	Context *gorm.DB
}

func JWTAuthService() UserAuthPermission {
	return &dbServices{
		Context: config.DB,
	}
}

func (db *dbServices) FindbyUserid(userId int) entitys.UserAuthPermission {
	user_permission := entitys.UserAuthPermission{}
	db.Context.First(&user_permission, userId)
	return user_permission
}

func (db *dbServices) Create(Data entitys.UserAuthPermission) {
	db.Context.Create(&Data)
}

func (db *dbServices) UpdateVersionToken(versionToken int, Data entitys.UserAuthPermission) {
	db.Context.Model(&Data).Where("user_id", Data.UserId).Update("token_version", versionToken).Update("iat", Data.Iat).Update("ext", Data.Ext)
}
