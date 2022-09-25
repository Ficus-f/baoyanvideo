package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type VideoListService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=200"`
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
