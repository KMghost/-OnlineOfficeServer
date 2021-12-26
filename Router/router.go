package Router

import (
	"OnlineOfficeServer/Controllers"
	"OnlineOfficeServer/Middlewares"
	"OnlineOfficeServer/Utils/VerficationCode"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用 「 跨域中间件 」， 否则 OPTIONS 会返回 404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	// router.Use(sessions.Sessions("myyyyyyysession", Sessions.Store))
	api := router.Group("api")
	{
		// api.POST("/testinsert", Controllers.TestInsert)
		api.GET("/captcha", VerficationCode.GetImg)
		api.GET("/system/config/menu", Controllers.MenuSelect)
		api.POST("/login", Controllers.Login)
		api.GET("/admin/info", Controllers.AdminInfo)
		api.GET("/system/basic/pos", Controllers.PositionSelect)
		api.POST("/system/basic/pos", Controllers.PositionInsert)
	}

	router.Run(":8081")
}
