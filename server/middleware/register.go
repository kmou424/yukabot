package middleware

import G "github.com/kmou424/yukabot/global"

func Register() {
	G.Router.Use(Auth())
}
