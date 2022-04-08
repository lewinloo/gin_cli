package service

import (
	"errors"
	"gin_cli/global"
	"gin_cli/model"
	"gin_cli/service/request"

	"gorm.io/gorm"
)

type ApiService struct {
}

// 新增 Api
func (s *ApiService) CreateApi(params request.CreateApiParams) (err error) {
	var api model.Api
	if !errors.Is(global.DB.Where("path = ? AND method = ?", params.Path, params.Method).First(&model.Api{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	api.ApiGroup = params.ApiGroup
	api.Method = params.Method
	api.Description = params.Description
	api.Path = params.Path
	return global.DB.Create(&api).Error
}

// 查询api
func (s *ApiService) GetApiList(params request.QueryApiListParams) (list interface{}, total int64, cur int, size int, err error) {
	if params.Current == 0 {
		params.Current = 1
	}

	if params.Size == 0 {
		params.Size = 10
	}
	limit := params.Current
	offset := params.Size * (params.Current - 1)
	db := global.DB.Model(&model.Api{})
	var apiList []model.Api

	if params.Path != "" {
		db = db.Where("path LIKE ?", "%"+params.Path+"%")
	}

	if params.Description != "" {
		db = db.Where("description LIKE ?", "%"+params.Description+"%")
	}

	if params.Method != "" {
		db = db.Where("method = ?", params.Method)
	}

	if params.ApiGroup != "" {
		db = db.Where("api_group = ?", params.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, params.Current, params.Size, err
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("api_group").Find(&apiList).Error
	}

	return apiList, total, params.Current, params.Size, err
}

// 获取所有的api
func (s *ApiService) GetAllApis() (apis []model.Api, err error) {
	err = global.DB.Find(&apis).Error
	return apis, err
}
