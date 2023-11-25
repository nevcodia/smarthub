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

	group.GET("/:type/")
}
