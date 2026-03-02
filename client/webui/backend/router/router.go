package router

import (
	"gostc-sub/webui/backend/bootstrap"

	"github.com/gin-gonic/gin"
)

var account string
var password string

func SetBasicAuth(u, p string) {
	account = u
	password = p
}

func init() {
	bootstrap.Route = func(engine *gin.Engine) {
		if account != "" && password != "" {
			engine.Use(gin.BasicAuth(gin.Accounts{
				account: password,
			}))
		}
		InitStatic(engine)
		InitApi(engine)
	}
}
