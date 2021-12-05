package sendmsg

import (
	"context"
	"encoding/base64"

	"github.com/fox-one/mixin-sdk-go"
)


func SendTextMsg(ctx context.Context,c *mixin.Client, text, conversation_id, recipient_id string) error {
	vmr := genTextMsg(text, conversation_id, recipient_id)
	return c.SendMessage(ctx, vmr)
}

func genTextMsg(text, conversation_id, recipient_id string) *mixin.MessageRequest {
	data := genTextData(text)
	rq := genMsgRequest(conversation_id, recipient_id, mixin.MessageCategoryPlainText, data)
	return rq
}

func genTextData(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

