package sendmsg

import (
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
		MessageID:      id.String(),
		Category:       category,
		Data:           data,
	}
}