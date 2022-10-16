package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=300"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=600"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "创建视频失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
