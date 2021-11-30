package Models

import (
	"OnlineOfficeServer/Databases/Mysql"
)

type Test struct {
	Id      int
	Testcol string `gorm:"colum:testcol"`
}

// 设置 Test 的表名为 `test`
func (Test) TableName() string {
	return "test"
}

func (this *Test) Insert() (id int, err error) {
	result := Mysql.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
