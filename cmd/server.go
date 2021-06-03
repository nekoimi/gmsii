package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nekoimi/gmsii/bot"
	"github.com/nekoimi/gmsii/config"
	"github.com/nekoimi/gmsii/cron"
	"github.com/nekoimi/gmsii/message"
	"github.com/nekoimi/gmsii/router"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var (
	port       string
	bind       string
	confjson   string
	confglobal = config.GlobalConfig
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Example:gmsii server --port=8080",
	Example: "gmsii server --bind=0.0.0.0 --port=8080 -c config.json",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	flags := serverCmd.PersistentFlags()
	flags.StringVarP(&port, "port", "p", "8000", "listen port")
	flags.StringVarP(&bind, "bind", "b", "127.0.0.1", "bind address")
	flags.StringVarP(&confjson, "config", "c", "config.json", "specify configuration file")

	go singleBroadcaster()
	go bootCronScheduler()
}

func run() {
	confglobal.ParseJson(confjson)

	if !confglobal.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	router.InitApiRouter(engine)

	listenAddress := bind + ":" + port
	log.Println("start...\tserver listen to http://" + listenAddress)
	if err := engine.Run(listenAddress); err != nil {
		log.Fatalln(err)
	}
}

func singleBroadcaster() {
	for {
		msg := <-message.Pipeline
		for _, messageSender := range bot.Senders {
			if err := messageSender.Send(msg); err != nil {
				fmt.Println(err)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func bootCronScheduler() {
	cron.Scheduler.Start()
}
