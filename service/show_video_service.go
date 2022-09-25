package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type ShowVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=200"`
}

// Show 创建视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "查询视频失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
