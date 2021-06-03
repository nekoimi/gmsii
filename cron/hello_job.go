package cron

import (
	"encoding/json"
	"fmt"
	"github.com/nekoimi/gmsii/config"
	"github.com/nekoimi/gmsii/message"
	"github.com/nekoimi/gmsii/utils"
	"time"
)

const (
	requestApi = "https://v1.hitokoto.cn/?c=a&encode=json&charset=utf-8"
)

type Hello struct{}

type Hitokoto struct {
	Content string `json:"hitokoto"`
	From    string `json:"from"`
}

func (h *Hello) Run() {
	if !config.GlobalConfig.EnableHeartbeat {
		return // ignore
	}
	response, err := utils.Get(requestApi, nil)
	if err != nil {
		fmt.Println(err)
	}
	result := Hitokoto{}
	_ = json.Unmarshal(response, &result)

	go func() {
		message.Pipeline <- message.NewMarkdown(fmt.Sprintf(`
<font color="info">%s</font>
> <font color="comment">%s</font> —— 「<font color="comment">%s</font>」
`, time.Now().Format("2006-01-02 15:04:05"), result.Content, result.From))
	}()
}
