package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Env struct {
	EnvType    string
	SERVERport string
	DBusername string 
	DBname     string 
	DBmode     string  
	DBhost     string
	DBport     string
	DBpassword string
	SECRET_KEY string
}

func NewEnv() (*Env, error) {
	env := Env{}
	
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		return &env, err
	}

	if err := godotenv.Load(); err != nil {
		return &env, err
	}

	env.EnvType = viper.GetString("Enviroment")
	env.DBusername = viper.GetString("db.Username")
	env.DBname = viper.GetString("db.Name")
	env.DBmode = viper.GetString("db.Mode")
	env.DBhost = viper.GetString("db.Host")
	env.DBport = viper.GetString("db.Port")
	
	env.DBpassword = os.Getenv("DB_PASSWORD")
	env.SECRET_KEY = os.Getenv("SECRET_KEY")  
	env.SERVERport = viper.GetString("server.Port")
	return &env, nil
}
