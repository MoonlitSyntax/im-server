package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Msg struct {
	Type         int8          `json:"type"`    //和MsgType一样
	Content      *string       `json:"content"` // 为1
	ImageMsg     *ImageMsg     `json:"imageMsg"`
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"`
	VideoMsg     *VideoMsg     `json:"videoMsg"`
	FileMsg      *FileMsg      `json:"fileMsg"`
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"`
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`
	ReplyMsg     *ReplyMsg     `json:"replyMsg"`
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`
	AtMsg        *AtMsg        `json:"atMsg"`
}

func (c *Msg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

func (c Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type ImageMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}

type VideoMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 单位秒
}
type FileMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"`
	Type  string `json:"type"`
}
type VoiceMsg struct {
	Time int    `json:"time"` // 单位秒
	Src  string `json:"src"`
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	EndReason int8      `json:"endReason"` // 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}
type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	EndReason int8      `json:"endReason"` // 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通

}
type WithdrawMsg struct {
	Content   string `json:"content"`
	OriginMsg *Msg   `json:"originMsg"`
}
type ReplyMsg struct {
	MsgID   uint   `json:"msgID"`
	Content string `json:"content"`
	Msg     *Msg   `json:"msg"`
}
type QuoteMsg struct {
	MsgID   uint   `json:"msgID"`
	Content string `json:"content"`
	Msg     *Msg   `json:"msg"`
}

type AtMsg struct {
	UserID  uint   `json:"userID"`
	Content string `json:"content"`
	Msg     *Msg   `json:"msg"`
}
