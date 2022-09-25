package model

import "gorm.io/gorm"

// Video 用户模型
type Video struct {
	gorm.Model
	Title string
	Info  string
}
