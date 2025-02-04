package models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                        `json:"sendUserID"`
	SendUserModel        UserModel                   `gorm:"foreignKey:SendUserID" json:"-"`
	RevUserID            uint                        `json:"revUserID"`
	RevUserModel         UserModel                   `gorm:"foreignKey:RevUserID" json:"-"`
	Status               int8                        `json:"status"`                            //1 同意  2 拒绝 3 忽略 4 是否展示
	AdditionalMessage    *string                     `gorm:"size:128" json:"additionalMessage"` // 验证消息
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"`
}
