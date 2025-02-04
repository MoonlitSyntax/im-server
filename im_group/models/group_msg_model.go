package models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

type GroupMsgModel struct {
	models.Model
	GroupID    uint       `json:"groupID"`
	GroupModel GroupModel `gorm:"foreignKey:GroupID" json:"-"`
	SendUserID uint       `json:"sendUserID"`
	RevUserID  uint       `json:"revUserID"`
	MsgType    int8       `json:"msgType"` //消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语言 7 视频 8 撤回 9 回复 10 引用 11 at
	MsgPreview string     `gorm:"size:64" json:"msgPreview"`
	Msg        ctype.Msg  `json:"msg"`
}
