package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
	"os"
)

// Video 用户模型
type Video struct {
	gorm.Model
	Title   string
	Alias   string
	Episode int64
	Info    string
	URL     string
}

// 给视频的URL签名
func (video *Video) TokenURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.URL, oss.HTTPGet, 600)
	return signedGetURL
}
