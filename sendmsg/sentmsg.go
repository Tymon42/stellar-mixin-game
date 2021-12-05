package sendmsg

import (
	"stellar-mixin-game/db"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
)

func genMsgRequest(conversation_id, recipient_id, category, data string) *mixin.MessageRequest {
	// The incoming message's message ID, which is an UUID.
	// id, _ := uuid.FromString(msg_id)
	id, _ := uuid.NewV1()
	return &mixin.MessageRequest{
		ConversationID: conversation_id,
		RecipientID:    recipient_id,
		// MessageID:      uuid.NewV5(id, "reply").String(),
		MessageID: id.String(),
		Category:  category,
		Data:      data,
	}
}

func GenMsgRequest(conversation_id, recipient_id string, event *db.Event) *mixin.MessageRequest {
	switch event.EventCategory {
	case mixin.MessageCategoryPlainText:
		return genTextMsg(event.Event,conversation_id,recipient_id)
	case mixin.MessageCategoryPlainVideo:
		return genVideoMsg(event.Event,conversation_id,recipient_id)
	}
	return &mixin.MessageRequest{}
}
func GenMsgRequests(conversation_id, recipient_id string, events []*db.Event) []*mixin.MessageRequest {
	var requests []*mixin.MessageRequest
	for _, event := range events{
		r := GenMsgRequest(conversation_id, recipient_id, event)
		requests = append(requests, r)
	}
	return requests
}
