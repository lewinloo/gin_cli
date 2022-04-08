package service

import (
	"errors"
	"gin_cli/global"
	"gin_cli/model"
	"gin_cli/service/request"

	"gorm.io/gorm"
)

type RoleService struct {
}

// 创建一个角色
func (s *RoleService) CreateRole(params request.CreateRoleParams) (role model.Role, err error) {
	if !errors.Is(global.DB.Where("role_id = ?", params.RoleId).First(&role).Error, gorm.ErrRecordNotFound) {
		return role, errors.New("存在相同角色id")
	}
	role.RoleId = params.RoleId
	role.RoleName = params.RoleName
	role.ParentId = params.ParentId
	role.DefaultRouter = params.DefaultRouter
	err = global.DB.Create(&role).Error
	return role, err
}
