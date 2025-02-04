package models

import "im-server/common/models"

type FriendModel struct {
	models.Model
	SendUserID    uint      `json:"sendUserID"`
	RevUserID     uint      `json:"revUserID"`
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	Notice        string    `gorm:"size:128" json:"notice"` //备注
}
