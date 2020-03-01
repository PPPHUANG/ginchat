// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-29
// Time: 16:58

package chat

import (
	"ginchat/db_conn"
	"ginchat/model"

	log "github.com/sirupsen/logrus"
)

func SaveMessageUserToUser(message *model.MessageUserToUser) (int64, error) {
	_, err := db_conn.DbClient.InsertOne(message)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/httphandlerChat.go/saveMessageUserToUser/InsertOne",
			"data":     message,
		}).Error(err.Error())
	}
	return message.Id, err
}

func SaveCommunityMessage(message *model.CommunityMessage) (int64, error) {
	_, err := db_conn.DbClient.InsertOne(message)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/httphandlerChat.go/saveCommunityMessage/InsertOne",
			"data":     message,
		}).Error(err.Error())
	}
	return message.Id, err
}

func UpdateMessageUserToUserState(id int64, message *model.MessageUserToUser) error {
	_, err := db_conn.DbClient.Id(id).Update(message)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/httphandlerChat.go/UpdateMessageUserToUserState/Update",
			"id":       id,
			"data":     message,
		}).Error(err.Error())
	}
	return err
}

func UpdateCommunityMessageState(id int64, message *model.CommunityMessage) error {
	_, err := db_conn.DbClient.Update(message)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/httphandlerChat.go/UpdateCommunityMessageState/Update",
			"id":       id,
			"data":     message,
		}).Error(err.Error())
	}
	return err
}

func SaveCommunityMessageToUser(message *model.CommunityMessageToUser) (int64, error) {
	_, err := db_conn.DbClient.InsertOne(message)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/httphandlerChat.go/SaveCommunityMessageToUser/InsertOne",
			"data":     message,
		}).Error(err.Error())
	}
	return message.Id, err
}
