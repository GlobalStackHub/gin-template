package routers

import (
	//"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	//首页
	app.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "index.html") })
	apiPrefix := "/api"
	checkUserType(apiPrefix)

	api := app.Group(apiPrefix)
	RegisterApiRouter(api)
}

func checkUserType(apiPrefix string) {
	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/user/login")
	//g.Use(middleware.UserAuthMiddleware(
	//	middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	//))
	// 权限验证
	var notCheckPermissionUrlArr []string
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, notCheckLoginUrlArr...)
	notCheckPermissionUrlArr = append(notCheckPermissionUrlArr, apiPrefix+"/menu/menubuttonlist")
	//g.Use(middleware.CasbinMiddleware(
	//	middleware.AllowPathPrefixSkipper(notCheckPermissionUrlArr...),
	//))
	//sys

}
