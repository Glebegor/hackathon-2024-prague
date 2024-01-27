package route

import (
	"time"
	
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	bootstrap "ton-event-bot/bootstrap"

)

func NewEventsRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	group.GET("/")
}

