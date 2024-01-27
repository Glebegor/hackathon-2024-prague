package main

import (
	"time"

	"ton-event-bot/api/route"
	"ton-event-bot/bootstrap"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Postgres

	gin := gin.Default()

	timeout := 10 * time.Second

	route.Setup(env, timeout, db, gin)
	gin.Run(":" + env.SERVERport)
}
