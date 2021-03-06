package application

import (
	models "app-test/src/pkg"
	"app-test/src/pkg/config"
	"app-test/src/pkg/datasource"
	"app-test/src/pkg/logger"
	AppConfig "app-test/src/web/config"
	"app-test/src/web/entity"
	"app-test/src/web/middleware"
	routers "app-test/src/web/router"
	"app-test/support/convert"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

// 运行
func Run(configPath string) {
	// 初始化配置项
	if configPath == "" {
		configPath = "./config.yaml"
	}
	// 加载配置
	loadConfig, err := AppConfig.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	logger.InitLog("debug", loadConfig.Web.LogPath)

	initDB(loadConfig)
	logger.Debug("数据库加载完成.......")

	// 如果不存在则创建数据库
	if !datasource.DB.HasTable(&entity.Email{}) {
		datasource.DB.CreateTable(&entity.Email{})
	}

	initWeb(loadConfig)
}

func initDB(config *config.Config) {
	models.InitDB(config)
}

func initWeb(loadConfig *config.Config) {
	gin.SetMode(gin.DebugMode)
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	app.NoMethod(middleware.NoMethodHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())

	// 注册路由
	routers.RegisterRouter(app)
	go initHTTPServer(loadConfig, app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}
