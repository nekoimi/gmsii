package bot

import "github.com/nekoimi/gmsii/message"

var (
	Senders map[string]MessageSender
)

type MessageSender interface {
	Send(message *message.Message) error
}

func init() {
	Senders = make(map[string]MessageSender)
}

func registerSender()  {
	Senders["wechatSender"] = &WeChatSender{}
}
