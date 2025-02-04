package models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

type GroupModel struct {
	models.Model
	Title                string                      `gorm:"size:32" json:"title"`
	Abstract             string                      `gorm:"size:128" json:"abstract"`
	Avatar               string                      `gorm:"size:256" json:"avatar"`
	Creator              uint                        `json:"creator"`
	IsSearch             bool                        `json:"isSearch"`
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"`
	IsInvite             bool                        `json:"isInvite"`           //是否可邀请好友
	IsTemporarySession   bool                        `json:"isTemporarySession"` //是否开启临时会话
	IsProhibition        bool                        `json:"isProhibition"`      //是否开启全员禁言
	Size                 int                         `json:"size"`               // 群规模
}
