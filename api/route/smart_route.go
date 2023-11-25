package route

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/nevcodia/smarthub/api/controller"
	"github.com/nevcodia/smarthub/domain"
	"github.com/nevcodia/smarthub/repository"
	"github.com/nevcodia/smarthub/service"
)

func NewSmartRouter(client *s3.Client, group *gin.RouterGroup) {
	s3Repository := repository.NewS3Repository(client)
	reps := map[domain.StorageType]domain.StorageRepository{
		domain.S3: s3Repository,
	}
	smartController := controller.NewSmartController(service.NewSmartService(&reps))

	group.GET("/support", smartController.StorageTypes)
	group.GET("/:type/stores", smartController.StoreNames)
	group.GET("/:type/objects", smartController.Objects)
	group.GET("/:type/objects/detail", smartController.ObjectsWithMetadata)
	group.GET("/:type/object", smartController.GetObject)
	group.DELETE("/:type/object", smartController.Delete)
	group.DELETE("/:type/object/all", smartController.DeleteAll)
	group.POST("/:type/upload", smartController.Upload)
	group.GET("/:type/upload-link", smartController.PresignUploadLink)
	//group.POST("/:type/upload-link", smartController.PresignUploadLinkWithMetadata)
	group.GET("/:type/download-link", smartController.PresignDownloadLink)
	group.GET("/:type/download", smartController.Download)
	group.PUT("/:type/copy", smartController.Copy)
	//group.PUT("/:type/copy/multi", smartController.CopyMulti)
	group.PUT("/:type/copy/all", smartController.CopyAll)
	group.PUT("/:type/move", smartController.Move)

}
