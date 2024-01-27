package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"ton-event-bot/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("api/v1/")		
 	NewEventsRouter(env, timeout, db, publicRouter)
	
	// protectRouter := gin.Group("api/v2/")
	// protectRouter.Use(middleware.Auth(env.SECRET_KEY))
}
