package main

import (
	// "bytes"
	"context"
	"encoding/base64"

	// "encoding/base64"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
	// "os"

	"os/signal"
	"syscall"
	"time"

	"stellar-mixin-game/db"
	"stellar-mixin-game/sendMsg"
	"stellar-mixin-game/util"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
)

func main() {
	store, err := util.StartMixin("./keystore.json")
	util.CheckErr(err)
	client, err := mixin.NewFromKeystore(&store.Store)
	util.CheckErr(err)
	// fmt.Printf("client: %+v\n", client)

	ctx := context.Background()

	fv := "d384d0a0-78d7-4f14-9761-21e85586a18a"

	sdb := db.OpenDB()

	sdb.AutoMigrate(&db.Event{}, &db.User{})

	h := func(ctx context.Context, msg *mixin.MessageView, userID string) error {
		// if there is no valid user id in the message, drop it
		if userID, _ := uuid.FromString(msg.UserID); userID == uuid.Nil {
			return nil
		}

		//解析内容
		context_byte, _ := base64.StdEncoding.DecodeString(msg.Data)
		context := string(context_byte)

		switch context {
		case "你好":
			if !db.If_old_user(sdb, msg.UserID) {
				db.Insert_user(sdb, msg.UserID)
				//TODO: 发送0000
			} else {
				//TODO: 发送菜单
			}
		case "/回滚":
			{
				//创建事件链
				events := []*db.Event{}
				events = append(events, db.FindLastEvent(sdb, msg.UserID))
				for events[len(events)-1].IsStop != true {
					events = append(events, db.FindLastEvent(sdb, msg.UserID))
				}
				db.RefreshLastEventID(sdb, msg.UserID, events[len(events)-1].EventID)
			}
		default:
			//TODO: 根据发来 event_id 新建事件链
		}

		return sendmsg.SendVideoMsg(ctx, client, fv, msg.ConversationID, msg.UserID)

		//TODO: 生成消息队列

		// Send the response
		// return client.SendMessage(ctx, reply)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// Start the message loop.
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			if err := client.LoopBlaze(ctx, mixin.BlazeListenFunc(h)); err != nil {
				log.Printf("LoopBlaze: %v", err)
			}
		}
	}

}
