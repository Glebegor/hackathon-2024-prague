package bootstrap

import (
	"log"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Env *Env
	Postgres *sqlx.DB
}

func App() Application {
	app := &Application{}
	env, err := NewEnv()
	if err != nil {
		log.Fatalln(err)
	}
	app.Env = env
	db, err := NewPostgresConnection(env)
	if err != nil {
		log.Fatalln(err)
	}
	app.Postgres = db
	return *app
}
