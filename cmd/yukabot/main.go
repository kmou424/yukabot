package main

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/sysutil"
	"github.com/kmou424/dingtalkbot"
	"github.com/kmou424/yukabot/bot"
	_ "github.com/kmou424/yukabot/cmd/cli"
	G "github.com/kmou424/yukabot/global"
	. "github.com/kmou424/yukabot/global/logger"
	"github.com/kmou424/yukabot/server"
)

func init() {
	var err error

	G.BotClient, err = dingtalkbot.NewClient(
		os.Getenv("DINGTALK_BOT_CLIENT_ID"),
		os.Getenv("DINGTALK_BOT_CLIENT_SECRET"),
	)
	if err != nil {
		Logger.Error("create bot failed", "err", err)
		os.Exit(1)
	}

	bot.Init()
	server.Init()
}

func startBot() {
	err := G.BotClient.
		AutoReconnect().
		Debug(false).
		Register(dingtalkbot.TypeChat, G.BotModuleChatChain).
		Start()
	if err != nil {
		Logger.Error("start bot failed", "err", err)
		os.Exit(1)
	}
}

func main() {
	go startBot()

	host := sysutil.Getenv("SERVER_HOST", "0.0.0.0")
	port := sysutil.Getenv("SERVER_PORT", "8080")

	err := G.Router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		Logger.Error("start server failed", "err", err)
	}
}
