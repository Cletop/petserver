package routes

import (
	"github.com/chagspace/petserver/common"
	"github.com/chagspace/petserver/controller"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(file_router *gin.RouterGroup) {
	// init tencent oss client
	common.InitOSSClient()

	file_router.GET("/file/:id", controller.GetFile)
	file_router.POST("/file", controller.UploadFile)
	file_router.DELETE("/file/:id", controller.DeleteFile)
	file_router.PUT("/file/:id/visibility", controller.SetVisibility)
	file_router.GET("/file/:id/download", controller.DownloadFile)

	//
	file_router.GET("/file/:id/thumbnail", controller.GetThumbnail)
	file_router.GET("/file/:id/thumbnail/:width/:height", controller.GetThumbnail)

	// get buckets
	file_router.GET("/buckets", controller.GetBuckets)
}
