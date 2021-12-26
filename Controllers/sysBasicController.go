package Controllers

import (
	"OnlineOfficeServer/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 职称管理
func PositionSelect(c *gin.Context) {
	var position Services.Position
	result := position.Select()
	c.JSON(http.StatusOK, result)
}

func PositionInsert(c *gin.Context) {
	json := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&json)
	name := json["name"].(string)
	var position Services.Position
	position.Insert(name)
	c.JSON(http.StatusOK, gin.H{
		"code": 200, "message": "添加成功",
	})
}

// 职位管理
func JobLevelSelect(c *gin.Context) {
	var jobLevel Services.JobLevel
	result := jobLevel.Select()
	c.JSON(http.StatusOK, result)
}
