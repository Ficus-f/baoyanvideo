package serializer

import "baoyanvideo/model"

// Video 视频序列化器
type Video struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Info  string `json:"info"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildVideoResponse 序列化视频响应
func BuildVideoResponse(video model.Video) Response {
	return Response{
		Data: BuildVideo(video),
	}
}
