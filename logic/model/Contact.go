package models

import "github.com/freezeChen/jsontime"

type Contact struct {
	OwnerUid   int64             `json:"owner_uid" xorm:"'owner_uid'"`     //null
	OtherUid   int64             `json:"other_uid" xorm:"'other_uid'"`     //null
	Mid        int64             `json:"mid" xorm:"'mid'"`                 //null
	Type       int64             `json:"type" xorm:"'type'"`               //null
	CreateTime jsontime.JsonTime `json:"create_time" xorm:"'create_time'"` //null
}

func (Contact) TableName() string {
	return "contact"
}

type ContactVo struct {
	Uid      int64
	UserName string
	Mid      int64
	Content  string
	Online   bool
	Time     jsontime.JsonTime
}
