package route

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/nevcodia/smarthub/bootstrap"
)

func Setup(env *bootstrap.Env, s3Client *s3.Client, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	NewSmartRouter(s3Client, publicRouter)
}
