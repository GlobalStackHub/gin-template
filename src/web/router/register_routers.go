package routers

import (
	"app-test/src/web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(app *gin.RouterGroup) {
	testController := controller.TestController{}

	app.GET("/test/insert", testController.Insert)

	app.GET("/test/other", testController.Other)
}
