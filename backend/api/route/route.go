package route

import (
	"time"

	"ton-event-bot/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, gin *gin.Engine) {
	publicRouter := gin.Group("api/v1/")
	NewSubRouter(env, timeout, db, publicRouter)
	NewUsersubRouter(env, timeout, db, publicRouter)

	// protectRouter := gin.Group("api/v2/")
	// protectRouter.Use(middleware.Auth(env.SECRET_KEY))
}
