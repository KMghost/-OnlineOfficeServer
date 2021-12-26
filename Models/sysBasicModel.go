package Models

// 职位管理
type Position struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	CreateDate *MyTime `gorm:"<-:false" json:"createDate"`
	Enabled    bool    `gorm:"default:true" json:"enabled"`
}

func (Position) TableName() string { return "t_position" }

// 职称管理
type JobLevel struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	TitleLevel string  `json:"titleLevel"`
	CreateDate *MyTime `gorm:"<-:false" json:"createDate"`
	Enabled    bool    `gorm:"default:true" json:"enabled"`
}

func (JobLevel) TableName() string { return "t_joblevel" }
