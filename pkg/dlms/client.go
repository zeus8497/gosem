package dlms

import (
	"log"
	"time"
)

type Notification struct {
	ID               string
	DataNotification DataNotification
}

//go:generate mockery --name Client --structname ClientMock --filename clientMock.go

// Client specifies the client layer.
type Client interface {
	Connect() error
	Disconnect() error
	IsConnected() bool
	GetSettings() Settings
	SetSettings(settings Settings)
	SetAddress(client int, server int)
	SetLogger(logger *log.Logger)
	Associate() error
	CloseAssociation() error
	IsAssociated() bool
	SetNotificationChannel(id string, nc chan Notification)
	GetRequest(att *AttributeDescriptor, data interface{}) (err error)
	GetRequestWithSelectiveAccessByDate(att *AttributeDescriptor, start time.Time, end time.Time, data interface{}) (err error)
	GetRequestWithStructOfElements(data interface{}) (err error)
	SetRequest(att *AttributeDescriptor, data interface{}) (err error)
	SetRequestWithStructOfElements(data interface{}) (err error)
	ActionRequest(mth *MethodDescriptor, data interface{}) (err error)
	CheckRequestWithStructOfElements(data interface{}) (err error)
}
