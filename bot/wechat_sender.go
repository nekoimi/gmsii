package bot

import (
	"encoding/json"
	"errors"
	"github.com/nekoimi/gmsii/config"
	"github.com/nekoimi/gmsii/message"
	"github.com/nekoimi/gmsii/utils"
)

const (
	botWebHookUrl        = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
	maxSendMessageLength = 1024
)

var (
	confglobal = config.GlobalConfig
)

type WeChatSender struct{}

type WeChatResult struct {
	ErrCode    int32  `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

func (sender *WeChatSender) Send(msg *message.Message) error {
	url := botWebHookUrl + confglobal.BotHookKey

	params := make(map[string]interface{})
	content := make(map[string]interface{})

	params["msgtype"] = msg.MsgType

	runes := []rune(msg.Content)
	sLen := len(runes)
	if len(runes) >= maxSendMessageLength {
		sLen = maxSendMessageLength
	}

	content["content"] = string(runes[:sLen])
	switch msg.MsgType {
	case message.Text:
		params["text"] = content
		break
	case message.Markdown:
		params["markdown"] = content
		break
	}

	response, err := utils.Post(url, params)
	if err != nil {
		return err
	}
	result := WeChatResult{}
	_ = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		return errors.New(result.ErrMessage)
	}
	return nil
}
