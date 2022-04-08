package model

import "gin_cli/global"

type Menu struct {
	global.BaseModel
	ParentId       uint            `json:"parentId" gorm:"comment:父菜单id"`      // 父菜单id
	Name           string          `json:"name" gorm:"comment:路由名称"`           //路由名称
	Path           string          `json:"path" gorm:"comment:路由路径"`           //路由路径
	Hidden         bool            `json:"hidden" gorm:"comment:是否在菜单列表中隐藏"`   // 是否在菜单列表中隐藏
	Component      string          `json:"component" gorm:"comment:对应前端的文件路径"` // 对应前端的文件路径
	Sort           int             `json:"sort" gorm:"comment:排序标记"`           // 排序标记
	KeepAlive      bool            `json:"keepAlive" gorm:"comment:是否缓存"`      // 是否缓存
	DefaultMenu    bool            `json:"defaultMenu" gorm:"comment:是否是基础路由"` // 是否是基础路由
	Title          string          `json:"title" gorm:"comment:菜单标题"`          //菜单标题
	Icon           string          `json:"icon" gorm:"comment:菜单Icon"`         // 菜单Icon
	MenuParameters []MenuParameter `json:"menuParameters"`                     // 菜单参数
	Roles          []Role          `json:"roles" gorm:"many2many:menu_role"`
}

type MenuParameter struct {
	global.BaseModel
	Type   string `json:"type" gorm:"comment:地址栏参数是params还是query"` // 地址栏参数是params还是query
	Key    string `json:"key" gorm:"comment:地址栏参数的key"`            // 地址栏参数的key
	Value  string `json:"value" gorm:"comment:地址栏参数的value"`        // 地址栏参数的value
	MenuID uint   // 菜单id
}
