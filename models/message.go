package models

import (
	"time"
)

type Message struct {
	messageId        int
	creationDate     time.Time
	lastUpdateDate   time.Time
	deletionDate     time.Time
	senderUserId     int
	recipientUserId  int
	messageType      int
	groupChatId      int
	encryptedMessage []byte
}
