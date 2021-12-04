package sendmsg

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/fox-one/mixin-sdk-go"
)


func SendVideoMsg(ctx context.Context,c *mixin.Client, attachment_id, conversation_id, recipient_id string) error {
	vmr := genVideoMsg(attachment_id, conversation_id, recipient_id)
	return c.SendMessage(ctx, vmr)
}

func genVideoMsg(attachment_id, conversation_id, recipient_id string) *mixin.MessageRequest {
	data := genVideoData(attachment_id)
	rq := genMsgRequest(conversation_id, recipient_id, mixin.MessageCategoryPlainVideo, data)
	return rq
}


func SENT(ctx context.Context, client *mixin.Client, reply *mixin.MessageRequest) error  {
	// Send the response
	return client.SendMessage(ctx, reply)
}

func genVideoData(attachment_id string) string {
	msg_data := &mixin.VideoMessage{
		AttachmentID: attachment_id,
		MimeType:  "video/mp4",
		Width:     1024,
		Height:    1024,
		Size:      1024,
		Duration:  60,
	}
	msg_data_json, _ := json.Marshal(msg_data)
	return base64.StdEncoding.EncodeToString(msg_data_json)
}

