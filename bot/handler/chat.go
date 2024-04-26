package handler

import (
	"github.com/google/uuid"
	"github.com/kmou424/dingtalkbot"
	G "github.com/kmou424/yukabot/global"
	"github.com/kmou424/yukabot/global/data"
	"github.com/kmou424/yukabot/templates"
)

func Register() {
	G.BotModuleChatChain.
		Handle("id", IdHandler).
		Handle("help", HelpHandler).
		Handle("channel", ChannelHandler).
		Default(UnknownHandler)
}

var UnknownHandler dingtalkbot.HandlerFunc = func(c *dingtalkbot.Context) {
	c.Client.SendTextMessage(
		c.Chat().ConversationId,
		templates.String(
			templates.Unknown(),
		),
	)
}

var HelpHandler dingtalkbot.HandlerFunc = func(c *dingtalkbot.Context) {
	c.Client.SendTextMessage(
		c.Chat().ConversationId,
		templates.String(
			templates.Help(),
		),
	)
}

var IdHandler dingtalkbot.HandlerFunc = func(c *dingtalkbot.Context) {
	c.Client.SendMarkdownMessage(
		c.Chat().ConversationId,
		"ID信息",
		templates.Markdown(
			templates.ID(c.Chat()),
		),
	)
}

var ChannelHandler dingtalkbot.HandlerFunc = func(c *dingtalkbot.Context) {
	respText := func() string {
		if len(c.Args()) != 1 {
			return templates.String(templates.ChannelHelp())
		}
		switch c.Args()[0] {
		case "save":
			id := uuid.New().String()
			data.Data.Channels.Each(func(key, val string) bool {
				if val == c.Chat().ConversationId {
					id = ""
					return false
				}
				return true
			})
			if id != "" {
				data.Data.Channels.Put(id, c.Chat().ConversationId)
				data.SyncDataToFS()
			}
			return templates.String(templates.ChannelSave(id))
		case "view":
			var id string
			data.Data.Channels.Each(func(key, val string) bool {
				if val == c.Chat().ConversationId {
					id = key
					return false
				}
				return true
			})
			return templates.String(templates.ChannelView(id))
		case "delete":
			var id string
			data.Data.Channels.Each(func(key, val string) bool {
				if val == c.Chat().ConversationId {
					id = key
					return false
				}
				return true
			})
			var result bool
			if id != "" {
				data.Data.Channels.Delete(id)
				data.SyncDataToFS()
				result = true
			}
			return templates.String(templates.ChannelDelete(result))
		default:
			return templates.String(templates.ChannelHelp())
		}
	}()
	c.Client.SendTextMessage(
		c.Chat().ConversationId,
		respText,
	)
}
