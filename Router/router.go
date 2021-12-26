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
		// 获取验证码
		api.GET("/captcha", VerficationCode.GetImg)
		// 获取菜单列表
		api.GET("/system/config/menu", Controllers.MenuSelect)
		// 登录
		api.POST("/login", Controllers.Login)
		api.POST("/logout", Controllers.Logout)
		// 用户权限
		api.GET("/admin/info", Controllers.AdminInfo)
		// 职位管理
		api.GET("/system/basic/pos", Controllers.PositionSelect)
		api.POST("/system/basic/pos", Controllers.PositionInsert)
		// 职称管理
		api.GET("/system/basic/joblevel", Controllers.JobLevelSelect)
	}

	router.Run(":8081")
}
