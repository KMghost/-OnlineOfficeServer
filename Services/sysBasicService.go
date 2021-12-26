package Services

import (
	"OnlineOfficeServer/Databases/Mysql"
	"OnlineOfficeServer/Models"
)

type Position Models.Position

func (this Position) Select() (result []Models.Position) {
	Mysql.DB.Order("id").Find(&result)
	return
}

func (this Position) Insert(name string) error {
	var position Models.Position
	position.Name = name
	result := Mysql.DB.Omit("ID").Debug().Create(&position)
	return result.Error
}
