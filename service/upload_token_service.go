package service

import (
	"baoyanvideo/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// "github.com/google/uuid"
	"os"
	"strings"
)

// UploadTokenService 获得上传 oss token 服务
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

func (service *UploadTokenService) GetToken() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	// 获取存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "获取OSS空间错误",
			Error: err.Error(),
		}
	}

	// 带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType("video/mp4"),
	}
	alias_filename := strings.Split(service.Filename, "_")
	alias := alias_filename[0]
	filename := alias_filename[1]
	key := "video/" + alias + "/" + filename
	// 签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS Put 错误",
			Error: err.Error(),
		}
	}

	// signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	// if err != nil {
	// 	return serializer.Response{
	// 		Code:  50002,
	// 		Msg:   "OSS Get 错误",
	// 		Error: err.Error(),
	// 	}
	// }
	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			// "get": signedGetURL,
		},
	}
}
