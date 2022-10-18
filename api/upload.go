package api

import (
	"baoyanvideo/service"
	"github.com/gin-gonic/gin"
)

func UploadToken(c *gin.Context) {
	service := service.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetToken()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
