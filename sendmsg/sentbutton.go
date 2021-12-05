package sendmsg

import (
	"encoding/base64"
	"encoding/json"
	"stellar-mixin-game/db"

	"github.com/fox-one/mixin-sdk-go"
)

func GenButtonMsg(button_group mixin.AppButtonGroupMessage, conversation_id, recipient_id string) *mixin.MessageRequest {
	data := genButtonData(button_group)
	rq := genMsgRequest(conversation_id, recipient_id, mixin.MessageCategoryAppButtonGroup, data)
	return rq
}


func genButtonData(button_group mixin.AppButtonGroupMessage) string {
	msg_data_json, _ := json.Marshal(button_group)
	return base64.StdEncoding.EncodeToString(msg_data_json)
}

func Gen2ButtonData(levent, revent db.Event) mixin.AppButtonGroupMessage {
	button_group := mixin.AppButtonGroupMessage{}
	lmsg_data := &mixin.AppButtonMessage{
		Label: levent.EventTitle,
		Action: "input:"+levent.EventID,
		Color: "#2A92F1",
	}
	button_group = append(button_group, *lmsg_data)
	rmsg_data := &mixin.AppButtonMessage{
		Label: revent.EventTitle,
		Action: "input:"+revent.EventID,
		Color: "#540907",
	}
	button_group = append(button_group, *rmsg_data)
	return button_group
}