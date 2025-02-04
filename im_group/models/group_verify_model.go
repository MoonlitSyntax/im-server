package models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

type GroupVerifyModel struct {
	models.Model
	GroupID              uint                        `json:"groupID"`
	GroupModel           GroupModel                  `gorm:"foreignKey:GroupID" json:"-"`
	UserID               uint                        `json:"userID"`
	Status               int8                        `json:"status"`
	AdditionalMessage    string                      `gorm:"size:32" json:"additionalMessage"`
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"`
	Type                 int8                        `json:"type"`
}
