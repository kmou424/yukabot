package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kmou424/yukabot/app/model"
	G "github.com/kmou424/yukabot/global"
	"github.com/kmou424/yukabot/global/data"
	"github.com/kmou424/yukabot/templates"
)

var alarmText gin.HandlerFunc = func(c *gin.Context) {
	textAlarm := &model.TextAlarm{}
	c.ShouldBindJSON(textAlarm)

	channelId := textAlarm.Channel
	conversationId, ok := data.Data.Channels.Get(channelId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel not found"})
		return
	}

	if strings.Trim(textAlarm.Content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content is empty"})
		return
	}

	if int(textAlarm.Level) >= len(templates.LevelsPrefix) || textAlarm.Level < 0 {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "level is invalid"})
		return
	}

	G.BotClient.Messenger.SendTextMessage(
		conversationId,
		templates.String(templates.TextAlarm(textAlarm.Level, textAlarm.Content)),
	)

	c.JSON(http.StatusOK, gin.H{"message": "text alarm sent"})
}
