package server

import (
	"github.com/kmou424/yukabot/server/handler"
	"github.com/kmou424/yukabot/server/middleware"
)

func Init() {
	middleware.Register()
	handler.Register()
}
