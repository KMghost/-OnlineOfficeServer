package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"obj": map[string]string{
			"tokenHead": "afwdwe",
			"token":     "aefaw",
		},
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "撤销成功",
	})
}
