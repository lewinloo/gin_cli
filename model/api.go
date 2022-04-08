package model

import "gin_cli/global"

type Api struct {
	global.BaseModel
	Path        string `json:"path" gorm:"comment:api路径"`                // api路径
	Description string `json:"description" gorm:"comment:api描述"`         // api描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`             // api组
	Method      string `json:"method" gorm:"default:POST;comment:api路径"` // 请求方法：POST|GET|PUT|PATCH|DELETE
}
