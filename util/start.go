package util

import (
	"encoding/json"
	"os"

	"github.com/fox-one/mixin-sdk-go"
)

type MixinBot struct {
	Store mixin.Keystore `json:"store,omitempty"`
	Pin   string         `json:"pin,omitempty"`
	// Client_secret string         `json:"client_secret,omitempty"`
}

func StartMixin(path string) (*MixinBot, error) {
	// 读取配置文件
	f_keystore, err := os.ReadFile(path)
	CheckErr(err)

	var key struct {
		PIN        string `json:"pin"`
		ClientId   string `json:"client_id"`
		SessionId  string `json:"session_id"`
		PINToken   string `json:"pin_token"`
		PrivateKey string `json:"private_key"`
	}
	err = json.Unmarshal(f_keystore, &key)
	CheckErr(err)
	
	s := &mixin.Keystore{
		ClientID:   key.ClientId,
		SessionID:  key.SessionId,
		PrivateKey: key.PrivateKey,
		PinToken:   key.PINToken,
	}

	r := &MixinBot{
		Store: *s,
		Pin:   key.PIN,
	}


	return r, nil
}
