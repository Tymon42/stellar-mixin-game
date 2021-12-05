package db

import (
	"gorm.io/gorm"
)

type EventLink struct {
	Previous string
	LEvent   string
	REvent   string
}

type Event struct {
	gorm.Model
	EventLink     EventLink `gorm:"embedded"`
	EventCategory string
	EventTitle    string
	EventID       string
	Event         string
	IsSingal      bool
	IsStop        bool
}

type User struct {
	gorm.Model
	UserID      string `gorm:"unique"`
	LastEventID string //记录最后一个发生的首个事件的 EventID
}
