package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "没有该视频记录！",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "视频信息更新失败！",
			Error: err.Error(),
		}
	}
	return serializer.BuildVideoResponse(video)
}
