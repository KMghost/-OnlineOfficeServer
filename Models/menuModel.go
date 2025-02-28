package Models

// 菜单列表
type Menu struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	Path        string `json:"path"`
	Component   string `json:"component"`
	Name        string `json:"name"`
	IconCls     string `json:"iconCls" gorm:"column:iconCls"`
	KeepAlive   *bool  `json:"keepAlive" gorm:"column:keepAlive"`
	RequireAuth bool   `json:"requireAuth" gorm:"column:requireAuth"`
	Pid         int    `json:"pid"`
	Enabled     bool   `json:"enabled"`
}

func (Menu) TableName() string { return "t_menu" }
