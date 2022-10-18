package service

import (
	"baoyanvideo/model"
	"baoyanvideo/serializer"
)

type CreateVideoService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=300"`
	Alias   string `form:"alias" json:"alias"`
	Episode int64  `form:"episode" json:"episode"`
	Info    string `form:"info" json:"info" binding:"required,min=0,max=600"`
	URL     string `form:"url" json:"url"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title:   service.Title,
		Alias:   service.Alias,
		Episode: service.Episode,
		URL:     service.URL,
		Info:    service.Info,
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
