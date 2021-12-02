package Models

type Menu struct {
	Id          int
	Url         string
	Path        string
	Component   string
	Name        string
	Iconcls     string
	Keepalive   bool
	Requireauth bool
	Pid         int
	Enabled     bool
}

// 设置 Test 的表名为 `test`
func (Menu) TableName() string { return "menu" }
