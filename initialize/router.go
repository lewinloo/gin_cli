package initialize

import (
	"gin_cli/config"
	_ "gin_cli/docs"
	"gin_cli/global"
	"gin_cli/middleware"
	"gin_cli/routes"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routes() *gin.Engine {
	gin.SetMode(config.AppMode)
	Router := gin.Default()

	Router.Use(middleware.Cors(), GinLogger(global.LOG), GinRecovery(global.LOG, true))

	systemRouter := routes.RouterGroupApp

	if err := Router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 公共接口
	PublicGroup := Router.Group("")
	{
		systemRouter.InitPublicRoutes(PublicGroup)
	}

	// 需要登录才能访问的接口
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.CheckAuth(), middleware.CheckPermission())
	{
		systemRouter.InitTestRoutes(PrivateGroup)
		systemRouter.InitUserRoutes(PrivateGroup)
		systemRouter.InitRoleRoutes(PrivateGroup)
		systemRouter.InitApiRoutes(PrivateGroup)
		systemRouter.InitPermissionRoutes(PrivateGroup)
	}
	return Router
}
