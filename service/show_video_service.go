package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type ShowVideoService struct {
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

	return serializer.BuildShowVideoResponse(video)
}
