package models

import "github.com/freezeChen/jsontime"

type MessageRelation struct {
	OwnerUid   int64             `json:"owner_uid" xorm:"'owner_uid'"`     //null
	OtherUid   int64             `json:"other_uid" xorm:"'other_uid'"`     //null
	Mid        int64             `json:"mid" xorm:"'mid'"`                 //null
	Type       int64             `json:"type" xorm:"'type'"`               //null
	CreateTime jsontime.JsonTime `json:"create_time" xorm:"'create_time'"` //null
}

func (MessageRelation) TableName() string {
	return "message_relation"
}
