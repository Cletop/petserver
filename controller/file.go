package controller

import (
	"context"
	"fmt"

	"github.com/chagspace/petserver/common"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// TODO: implement tencent oss upload
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DeleteFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SetVisibility(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DownloadFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetBuckets(c *gin.Context) {
	services, _, err := common.GlobalOSSClient.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", services)
}

func GetThumbnail(c *gin.Context) {
	c.JSON(200, gin.H{})
}
