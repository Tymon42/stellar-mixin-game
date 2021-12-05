package main

import (
	// "bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"stellar-mixin-game/db"
	"stellar-mixin-game/sendmsg"
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

	// fv := "d384d0a0-78d7-4f14-9761-21e85586a18a"

	sdb := db.OpenDB("./game.db")

	sdb.AutoMigrate(&db.Event{}, &db.User{})

	h := func(ctx context.Context, msg *mixin.MessageView, userID string) error {
		// if there is no valid user id in the message, drop it
		if userID, _ := uuid.FromString(msg.UserID); userID == uuid.Nil {
			return nil
		}

		//解析内容
		context_byte, _ := base64.StdEncoding.DecodeString(msg.Data)
		context := string(context_byte)
		fmt.Printf("context: %+v\n", context)

		switch context {
		case "你好":
			fmt.Println("hi")
			if !db.If_old_user(sdb, msg.UserID) {
				db.Insert_user(sdb, msg.UserID)
				//TODO: 发送0000 按钮
				button_group := mixin.AppButtonGroupMessage{}
				lmsg_data := &mixin.AppButtonMessage{
					Label:  "开始游戏",
					Action: "input:" + "0000",
					Color:  "#2A92F1",
				}
				button_group = append(button_group, *lmsg_data)
				rq := sendmsg.GenButtonMsg(button_group,msg.ConversationID,msg.UserID)
				return client.SendMessage(ctx, rq)

			} else {
				//TODO: 发送菜单
				return sendmsg.SendTextMsg(ctx, client, "菜单", msg.ConversationID, msg.UserID)
			}
		case "/回滚":
			{
				//创建事件链
				events := []*db.Event{}
				events = append(events, db.FindLastEvent(sdb, msg.UserID))
				for events[len(events)-1].IsStop != true {
					//TODO: find next event
					events = append(events, db.FindLastEvent(sdb, msg.UserID))
				}
				db.RefreshLastEventID(sdb, msg.UserID, events[0].EventID)
				rq := sendmsg.GenMsgRequests(msg.ConversationID, msg.UserID, events)
				return client.SendMessages(ctx, rq)

			}
		default:
			{
				_, err := db.FindEvent(sdb, context)
				if err == nil {
					//根据发来 event_id 新建事件链
					events := []*db.Event{}
					event, _ := db.FindEvent(sdb, context)
					events = append(events, event)
					for events[len(events)-1].IsStop != true {
						events = append(events, db.QureyNextEvent(sdb, events[len(events)-1]))
					}
					if events[len(events)-1].IsStop {
						// 生成按钮
						levent := db.QureyLEvent(sdb, events[len(events)-1])
						revent := db.QureyREvent(sdb, events[len(events)-1])
						rq := sendmsg.GenButtonMsg(sendmsg.Gen2ButtonData(*levent, *revent), msg.ConversationID, msg.UserID)
						return client.SendMessage(ctx, rq)
					}
					db.RefreshLastEventID(sdb, msg.UserID, events[0].EventID)
					rq := sendmsg.GenMsgRequests(msg.ConversationID, msg.UserID, events)
					return client.SendMessages(ctx, rq)
				} else {
					return sendmsg.SendTextMsg(ctx, client, "nothing", msg.ConversationID, msg.UserID)
				}
			}
		}

		// return sendmsg.SendVideoMsg(ctx, client, fv, msg.ConversationID, msg.UserID)
		return sendmsg.SendTextMsg(ctx, client, "nothin", msg.ConversationID, msg.UserID)
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
