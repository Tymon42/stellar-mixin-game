package db

import (
	"fmt"
	"stellar-mixin-game/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	util.CheckErr(err)
	return db
}

func If_old_user(db *gorm.DB, user_id string) bool {
	var qUsers *User
	u := db.Where("user_id = ?", user_id).First(&qUsers)
	fmt.Printf("u: %+v\n", u)
	// fmt.Printf("u.statament: %+v\n",u.Statement)
	if u.Error != nil {
		return false
	} else {
		if qUsers.UserID == user_id {
			return true
		}
	}
	return false
}

func Insert_user(db *gorm.DB, user_id string) bool {
	user := User{
		UserID:      user_id,
		LastEventID: "0000",
	}
	db.Create(&user)
	return true
}

// 用于回滚事件
func FindLastEvent(db *gorm.DB, user_id string) *Event {
	var qUsers *User
	u := db.Where("user_id = ?", user_id).First(&qUsers)
	if u.Error != nil {
		return &Event{}
	} else {
		if qUsers.UserID == user_id {
			var nEvent *Event
			e := db.Where("event_id = ?", qUsers.LastEventID).First(&nEvent)
			if e.Error != nil {
				return nEvent
			} else {
				return nEvent
			}
		}
	}
	return &Event{}
}

func FindEvent(db *gorm.DB, event_id string) (*Event, error) {
	var event *Event
	u := db.Where("event_id = ?", event_id).First(&event)
	if u.Error != nil {
		return &Event{}, u.Error
	} else {
		if event.EventID == event_id {
			return event, nil
		}
	}
	return &Event{}, u.Error
}

func RefreshLastEventID(db *gorm.DB, user_id, last_event_id string) {
	db.Model(&User{}).Where("user_id = ?", user_id).Update("last_event_id", last_event_id)
}

func QureyNextEvent(db *gorm.DB, event *Event) *Event {
	var nEvent *Event
	u := db.Where("event_id = ?", event.LEvent).First(&nEvent)
	if u.Error != nil {
		return &Event{}
	} else {
		return nEvent
	}
	return &Event{}
}

func QureyLEvent(db *gorm.DB, event *Event) *Event {
	var nEvent *Event
	u := db.Where("event_id = ?", event.LEvent).First(&nEvent)
	if u.Error != nil {
		return &Event{}
	} else {
		return nEvent
	}
	return &Event{}
}

func QureyREvent(db *gorm.DB, event *Event) *Event {
	var nEvent *Event
	u := db.Where("event_id = ?", event.REvent).First(&nEvent)
	if u.Error != nil {
		return &Event{}
	} else {
		return nEvent
	}
	return &Event{}
}
