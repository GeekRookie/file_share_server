package dao

import (
	"file_share_server/global"
	"file_share_server/model"

	"gorm.io/gorm"
)

func GetAllShareFile() []*model.ShareFile {
	var list []*model.ShareFile
	global.DB.Find(&list)
	return list
}

func CreateShareFile(file *model.ShareFile) *gorm.DB {
	return global.DB.Create(file)
}

func DeleteShareFile(file *model.ShareFile) *gorm.DB {
	return global.DB.Delete(file)
}