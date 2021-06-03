package config

import (
	"encoding/json"
	"github.com/nekoimi/gmsii/utils"
	"io/ioutil"
	"log"
)

var (
	Version      = "1.0-dev"
	GlobalConfig *Config
)

type Config struct {
	Debug           bool   `json:"debug"`
	EnableHeartbeat bool   `json:"enable_heartbeat"`
	BotHookKey      string `json:"bot_hook_key"`
}

func init() {
	GlobalConfig = &Config{
		Debug: true,
	}
}

func (cfg *Config) ParseJson(jsonfile string) {
	if !utils.FileExists(jsonfile) {
		log.Fatalln(jsonfile + " does not exists!")
		return
	}
	jsonBytes, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		log.Fatalln("failed to read " + jsonfile + " configuration file!")
		return
	}
	_ = json.Unmarshal(jsonBytes, cfg)
}
