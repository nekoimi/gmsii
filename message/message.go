package message

import (
	"time"
)

var (
	Pipeline chan *Message
)

const (
	Text     = "text"
	Markdown = "markdown"
	Image    = "image"
)

type Message struct {
	Mid     int64
	MsgType string
	Content string
}

func init() {
	Pipeline = make(chan *Message)
}

func NewText(content string) *Message {
	return &Message{
		Mid:     time.Now().Unix(),
		MsgType: Text,
		Content: content,
	}
}

func NewMarkdown(content string) *Message  {
	return &Message{
		Mid:     time.Now().Unix(),
		MsgType: Markdown,
		Content: content,
	}
}
