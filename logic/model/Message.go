package models

import "github.com/freezeChen/jsontime"

type Message struct {
	Id          int64             `json:"id" xorm:"'id'"`                     //null
	Content     string            `json:"content" xorm:"'content'"`           //内容
	Type        int64             `json:"type" xorm:"'type'"`                 //null
	SenderId    int64             `json:"sender_id" xorm:"'sender_id'"`       //null
	RecipientId int64             `json:"recipient_id" xorm:"'recipient_id'"` //null
	CreateTime  jsontime.JsonTime `json:"create_time" xorm:"'create_time'"`   //null
}

func (Message) TableName() string {
	return "message"
}
