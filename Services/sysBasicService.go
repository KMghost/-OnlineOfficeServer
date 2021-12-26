package Services

import (
	"OnlineOfficeServer/Databases/Mysql"
	"OnlineOfficeServer/Models"
)

type Position Models.Position
type JobLevel Models.JobLevel

// 职称管理
func (this Position) Select() (result []Models.Position) {
	Mysql.DB.Order("id").Find(&result)
	return
}

func (this Position) Insert(name string) error {
	var position Position
	position.Name = name
	result := Mysql.DB.Omit("ID").Debug().Create(&position)
	return result.Error
}

// 职位管理
func (this JobLevel) Select() (result []Models.JobLevel) {
	Mysql.DB.Order("id").Find(&result)
	return
}

func (this JobLevel) Insert(name string) error {
	var position JobLevel
	position.Name = name
	result := Mysql.DB.Omit("ID").Debug().Create(&position)
	return result.Error
}
