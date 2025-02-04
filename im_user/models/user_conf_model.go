package models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

type UserConfModel struct {
	models.Model
	UserID               uint                        `json:"userID"`
	UserModel            UserModel                   `gorm:"foreignKey:UserID" json:"-"`
	RecallMessage        *string                     `gorm:"size:32" json:"recallMessage"` // 撤回消息的提示内容
	FriendOnline         bool                        `json:"friendOnline"`                 //上线提醒
	Sound                bool                        `json:"sound"`                        //是否静音
	SecureLink           bool                        `json:"secureLink"`                   //安全链接
	SavePwd              bool                        `json:"savePwd"`                      //记住密码
	SearchUser           int8                        `json:"searchUser"`                   //别人查找方法, 0 -> 不允许查找, 1 -> 通过用户号查找 2 -> 可以通过昵称查找
	FriendVerification   int8                        `json:"friendVerification"`           // 0-> 不允许任何人加好友, 1 -> 允许任何人 2-> 需要验证消息 3 -> 需要回答问题, 4 -> 需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"friendQuestion"`               // 验证问题
	Online               bool                        `json:"online"`
}
