package serializer

import "baoyanvideo/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Alias     string `json:"alias"`
	Episode   int64  `json:"episode"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		ID:        video.ID,
		Title:     video.Title,
		Info:      video.Info,
		URL:       video.TokenURL(),
		CreatedAt: video.CreatedAt.Unix(),
	}
}

// BuildShowVideoResponse 播放视频的响应
func BuildShowVideoResponse(video model.Video) Response {
	data := Video{
		ID:        video.ID,
		Title:     video.Title,
		Alias:     video.Alias,
		Episode:   video.Episode,
		Info:      video.Info,
		URL:       video.TokenURL(),
		CreatedAt: video.CreatedAt.Unix(),
	}
	return Response{
		Data: data,
	}
}

// BuildCreateVideoResponse 序列化简单视频后返回的信息响应
func BuildVideoResponse(video model.Video) Response {
	data := Video{
		ID:        video.ID,
		Title:     video.Title,
		Alias:     video.Alias,
		Episode:   video.Episode,
		Info:      video.Info,
		URL:       video.URL,
		CreatedAt: video.CreatedAt.Unix(),
	}
	return Response{
		Data: data,
	}
}

// BuildVideosResponse 序列化视频列表响应
func BuildVideosResponse(videos []model.Video) Response {
	var data []Video
	for _, item := range videos {
		video := Video{
			ID:        item.ID,
			Title:     item.Title,
			Info:      item.Info,
			CreatedAt: item.CreatedAt.Unix(),
		}
		data = append(data, video)
	}
	return Response{
		Data: data,
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := Video{
			ID:        item.ID,
			Title:     item.Title,
			Info:      item.Info,
			CreatedAt: item.CreatedAt.Unix(),
		}
		videos = append(videos, video)
	}
	return videos
}

// DataList 基础列表结构
type VideoList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

func BuildListVideosResponse(items interface{}, total uint) Response {
	return Response{
		Data: VideoList{
			Items: items,
			Total: total,
		},
	}
}
