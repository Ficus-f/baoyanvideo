package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type VideoListService struct {
}

func (service *VideoListService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "查询数据库错误",
			Error: err.Error(),
		}
	}
	return serializer.BuildVideosResponse(videos)
}
