package Models

import (
	"OnlineOfficeServer/Databases/Mysql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type Position struct {
	Id         int     `json:"id"`
	Name       string  `json:"name" `
	CreateDate *MyTime `gorm:"<-:false" json:"createDate"`
	Enabled    bool    `json:"enabled" `
}

func (Position) TableName() string { return "t_position" }

func (this Position) Select() []Position {
	var result []Position
	Mysql.DB.Order("id").Find(&result)
	return result
}

func (this Position) Insert(name string) error {
	this.Name = name
	this.Enabled = true
	result := Mysql.DB.Omit("ID").Debug().Create(&this)
	return result.Error
}

type MyTime time.Time

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

// 自定义类型上实现 Marshaler 的接口, 在进行 Marshal 时就会使用此除的实现来进行 json 编码
func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}
