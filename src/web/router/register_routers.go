package routers

import (
	"app-test/src/web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRouterSys(app *gin.RouterGroup) {
	menu := controller.Result{}
	app.GET("/test/insert", menu.Insert)
}
