package models

import (
	"time"
)

type Message struct {
	MessageId        int
	CreationDate     time.Time
	LastUpdateDate   time.Time
	DeletionDate     time.Time
	SenderUserId     int
	RecipientUserId  int
	MessageType      int
	GroupChatId      int
	EncryptedMessage string
}
