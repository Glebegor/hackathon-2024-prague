package bootstrap

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)



func NewPostgresConnection(env *Env) (*sqlx.DB, error){
	var db *sqlx.DB
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s", env.DBusername, env.DBname, env.DBmode, env.DBpassword, env.DBhost, env.DBport))
	if err != nil {
		return db, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return db, err
	} else {
		return db, nil
	}
}

