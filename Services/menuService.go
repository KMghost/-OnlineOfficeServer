package Services

import (
	"OnlineOfficeServer/Models"
)

type Menu struct {
	Id          int           `json:"id"`
	Url         string        `json:"url"`
	Path        string        `json:"path"`
	Component   string        `json:"component"`
	Name        string        `json:"name"`
	IconCls     string        `json:"iconCls"`
	KeepAlive   *bool         `json:"keepAlive"`
	RequireAuth *bool         `json:"requireAuth"`
	Pid         int           `json:"pid"`
	Enabled     bool          `json:"enabled"`
	Roles       *string       `json:"roles"`
	Children    []interface{} `json:"children"`
}

func (this *Menu) Select() []Menu {
	var menu Models.Menu
	records := menu.Select()
	var result []Menu
	var a Menu
	for _, data := range records {
		a.Id = data.Id
		a.Url = data.Url
		a.Path = data.Path
		a.Component = data.Component
		a.Name = data.Name

		a.IconCls = data.IconCls
		a.KeepAlive = data.KeepAlive
		a.RequireAuth = &data.RequireAuth
		a.Pid = data.Pid
		a.Enabled = data.Enabled
		if a.Children == nil {
			a.Children = make([]interface{}, 0)
		}
		// a.Roles = nil
		if a.Pid == 1 {
			result = append(result, a)
		} else if a.Pid > 1 {
			result[a.Pid-2].Children = append(result[a.Pid-2].Children, a)
		}

	}

	return result
}
