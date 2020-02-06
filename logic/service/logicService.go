package service

import (
	"DailysServer/logic/conf"
	"DailysServer/logic/dao"
	models "DailysServer/logic/model"
	"DailysServer/pkg/snowflake"
	"DailysServer/proto"

	"github.com/freezeChen/jsontime"
)

type LogicService struct {
	dao *dao.Dao
}

func NewLogicService(c *conf.Config, ) *LogicService {
	return &LogicService{dao: dao.New(c)}
}

func (svc *LogicService) SaveMessage(req *proto.MessageReq) error {
	now := jsontime.Now()
	var content = &models.Message{
		Id:          snowflake.GenID(),
		Content:     req.Content,
		Type:        req.Type,
		SenderId:    req.SenderId,
		RecipientId: req.RecipientId,
		CreateTime:  now,
	}
	session := svc.dao.NewSession()
	session.Begin()
	defer session.Close()

	_, err := session.InsertOne(content)
	if err != nil {
		return err
	}

	//写入消息消息索引表
	var relationSender = &models.MessageRelation{
		OwnerUid:   req.SenderId,
		OtherUid:   req.RecipientId,
		Mid:        content.Id,
		Type:       0,
		CreateTime: now,
	}
	_, err = session.InsertOne(relationSender)
	if err != nil {
		return err
	}

	var relationRecipient = &models.MessageRelation{
		OwnerUid:   req.RecipientId,
		OtherUid:   req.SenderId,
		Mid:        content.Id,
		Type:       1,
		CreateTime: now,
	}

	_, err = session.InsertOne(relationRecipient)
	if err != nil {
		return err
	}

	//更新发送者的最新联系人
	senderContact := svc.dao.GetContact(req.SenderId, req.RecipientId)
	if senderContact != nil {
		senderContact.Mid = content.Id
		senderContact.CreateTime = now

		_, err = session.Update(senderContact)
		if err != nil {
			return err
		}
	} else {
		senderContact = new(models.Contact)
		senderContact.OwnerUid = req.SenderId
		senderContact.OtherUid = req.RecipientId
		senderContact.Mid = content.Id
		senderContact.CreateTime = now
		senderContact.Type = 0
		_, err = session.InsertOne(senderContact)
		if err != nil {
			return err
		}
	}

	//更新接收者的最新联系人
	recipientContact := svc.dao.GetContact(req.RecipientId, req.SenderId)
	if recipientContact != nil {
		recipientContact.Mid = content.Id
		recipientContact.CreateTime = now

		_, err = session.Update(recipientContact)
		if err != nil {
			return err
		}
	} else {
		recipientContact = new(models.Contact)
		recipientContact.OwnerUid = req.RecipientId
		recipientContact.OtherUid = req.SenderId
		recipientContact.Mid = content.Id
		recipientContact.CreateTime = now
		recipientContact.Type = 1
		_, err = session.InsertOne(recipientContact)
		if err != nil {
			return err
		}
	}

	if err = session.Commit(); err != nil {
		return err
	}

	return nil
}

//获取联系人列表
func (svc *LogicService) GetContactList(uid int64) ([]*proto.Conversion, error) {
	return svc.dao.GetContactList(uid)
}

func (svc *LogicService) Online(sid string, uid int64) error {
	err := svc.dao.Online(sid, uid)
	if err != nil {
		return err
	}
	return err
}
