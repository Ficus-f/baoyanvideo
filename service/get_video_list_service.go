package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type VideoListService struct {
	Limit int    `form:"limit"`
	Start int    `form:"start"`
	Alias string `form:"alias"`
}

func (service *VideoListService) List() serializer.Response {
	var videos []model.Video
	var total int64 = 0

	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := model.DB.Model(model.Video{}).Where("alias = ?", service.Alias).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Where("alias = ?", service.Alias).Limit(service.Limit).Offset(service.Start).Order("episode desc").Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "查询数据库错误",
			Error: err.Error(),
		}
	}
	return serializer.BuildListVideosResponse(
		serializer.BuildVideos(videos),
		uint(total),
	)
}
