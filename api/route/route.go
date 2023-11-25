package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nevcodia/smarthub/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	// All Public APIs
	NewVideoRouter(db, publicRouter)
	NewUserRouter(db, publicRouter)

	uiRouter := gin.Group("/view")
	//User Interface
	NewUIRouter(db, uiRouter)
}
