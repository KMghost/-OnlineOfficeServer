package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":        1,
		"name":      "系统管理员",
		"phone":     "13812361398",
		"telephone": "71937538",
		"address":   "香港特别行政区强县长寿柳州路p座123",
		"enabled":   true,
		"username":  "admin",
		"password":  nil,
		"userFace":  "http://localhost:8080/static/touxiang.jpg",
		"remark":    nil,
		"roles": []map[string]interface{}{
			{
				"id":     6,
				"name":   "ROLE_admin",
				"nameZh": "系统管理员",
			},
		},

		"authorities": []map[string]interface{}{
			{
				"authority": "ROLE_admin",
			},
		},
		"accountNonExpired":     true,
		"credentialsNonExpired": true,
		"accountNonLocked":      true,
	})
}
