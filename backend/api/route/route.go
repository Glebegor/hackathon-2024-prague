package route

import (
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"ton-event-bot/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, gin *gin.Engine) {
	// publicRouter := gin.Group("")		
	fmt.Print("Something")
}
