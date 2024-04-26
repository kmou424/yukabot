package G

import (
	"github.com/gin-gonic/gin"
	"github.com/kmou424/dingtalkbot"
)

var (
	BotClient          *dingtalkbot.Client
	BotModuleChatChain             = dingtalkbot.ModuleChatChain()
	Router             *gin.Engine = gin.New()
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}
