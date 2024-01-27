package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"ton-event-bot/bootstrap"
	"ton-event-bot/api/route"

	_ "github.com/lib/pq"
)


func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Postgres
	
	gin := gin.Default()
	
	timeout := 10*time.Second

	route.Setup(env, timeout, db, gin)
	gin.Run(env.SERVERport)
}
