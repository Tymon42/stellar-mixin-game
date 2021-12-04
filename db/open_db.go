package db

import (
	"fmt"
	"log"
	"stellar-mixin-game/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
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
		log.Panicln(u.Error)
	} else {
		if qUsers.UserID == user_id {
			var nEvent *Event
			e := db.Where("event_id = ?", qUsers.LastEventID).First(&nEvent)
			if e.Error != nil {
				log.Panicln(e.Error)
			} else {
				return nEvent
			}
		}
	}
	return &Event{}
}

func RefreshLastEventID(db *gorm.DB, user_id, last_event_id string) {
	db.Model(&User{}).Where("user_id = ?", user_id).Update("last_event_id", last_event_id)
}
