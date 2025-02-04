package models

import "im-server/common/models"

type GroupMemberModel struct {
	models.Model
	GroupID         uint       `json:"groupID"`
	GroupModel      GroupModel `gorm:"foreignKey:GroupID" json:"-"`
	UserID          uint       `json:"userID"`
	MemberNickname  string     `gorm:"size:32" json:"memberNickname"`
	Role            int        `json:"role"`
	ProhibitionTime *int       `json:"prohibitionTime"` //单位分钟
}
