package main

import (
	"context"
	"fmt"
	"os"
	"stellar-mixin-game/util"

	"github.com/fox-one/mixin-sdk-go"
)

func main() {

	store, err := util.StartMixin("../keystore.json")
	util.CheckErr(err)
	client, err := mixin.NewFromKeystore(&store.Store)
	util.CheckErr(err)

	ctx := context.Background()

	files, _ := util.GetAllFiles("./", ".mp3")

	for _, file := range files {
		att, _ := client.CreateAttachment(ctx)

		f, _ := os.ReadFile(file)

		err = mixin.UploadAttachmentTo(ctx, att.UploadURL, f)
		util.CheckErr(err)
		fmt.Printf("att.AttachmentID: %v\n", att.AttachmentID)
	}
}
