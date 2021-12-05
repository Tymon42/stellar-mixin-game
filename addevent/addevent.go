package main

import (
	"fmt"
	"stellar-mixin-game/db"

	"github.com/fox-one/mixin-sdk-go"
	"gorm.io/gorm"
)

func main() {
	sdb := db.OpenDB("../game.db")
	sdb.AutoMigrate(&db.Event{}, &db.User{})

	event0 := db.Event{
		Previous:      "0000",
		LEvent:        "0001",
		REvent:        "0002",
		EventCategory: mixin.MessageCategoryPlainText,
		EventTitle:    "菜单",
		EventID:       "0000",
		Event:         "这是菜单",
		IsStop:        true,
	}

	r := InsertEvent(sdb, event0)
	fmt.Printf("r: %+v\n", r)

	event1 := db.Event{
		Previous:      "0000",
		LEvent:        "0003",
		REvent:        "0004",
		EventCategory: mixin.MessageCategoryPlainText,
		EventTitle:    "0001",
		EventID:       "0001",
		Event:         "这是0001",
		IsStop:        true,
	}

	r1 := InsertEvent(sdb, event1)
	fmt.Printf("r1: %+v\n", r1)

	event2 := db.Event{
		Previous:      "0000",
		LEvent:        "0005",
		REvent:        "0006",
		EventCategory: mixin.MessageCategoryPlainText,
		EventTitle:    "0002",
		EventID:       "0002",
		Event:         "这是0002",
		IsStop:        true,
	}

	r2 := InsertEvent(sdb, event2)
	fmt.Printf("r1: %+v\n", r2)
}

func InsertEvent(sdb *gorm.DB, event db.Event) *gorm.DB {
	return sdb.Create(&event)
}
