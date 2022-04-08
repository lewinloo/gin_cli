package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	RoleId        string         `gorm:"not null;unique;primaryKey;comment:角色id;size:90" json:"roleId"` // 角色id
	CreatedAt     time.Time      `json:"createdAt"`                                                     // 创建时间
	UpdatedAt     time.Time      `json:"updatedAt"`                                                     // 创建时间
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`                                                // 删除时间
	RoleName      string         `gorm:"comment:角色名" json:"roleName"`                                   //角色名
	ParentId      string         `json:"parentId" gorm:"comment:父角色ID"`                                 // 父角色ID
	DefaultRouter string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`           //默认菜单（默认为dashboard）
	Menus         []Menu         `json:"menus" gorm:"many2many:menu_role"`
	Users         []User         `json:"users"`
}
