package Controllers

import (
	"OnlineOfficeServer/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取菜单列表
func MenuSelect(c *gin.Context) {
	var menuService Services.Menu
	result := menuService.Select()
	c.JSON(http.StatusOK, result)
}
