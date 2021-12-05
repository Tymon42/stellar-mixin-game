package db

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Previous      string
	LEvent        string
	REvent        string
	EventCategory string
	EventTitle    string
	EventID       string `gorm:"unique"`
	Event         string
	IsStop        bool
}

type User struct {
	gorm.Model
	UserID      string `gorm:"unique"`
	LastEventID string //记录最后一个发生的首个事件的 EventID
}
