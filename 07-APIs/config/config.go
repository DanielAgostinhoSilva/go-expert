package config

import (
	"github.com/DanielAgostinhoSilva/go-expert/07-APIs/pkg/environment"
	"github.com/go-chi/jwtauth"
	"os"
)

var CFG Conf

type Conf struct {
	Database  Database
	WebServer WebServer
	Security  Security
}

type Database struct {
	DBDriver   string `env:"DB_DRIVER"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

type WebServer struct {
	Port int `env:"WEB_SERVER_PORT"`
}

type Security struct {
	JWTSecret   string `env:"JWT_SECRET"`
	JWtExpireIn int    `env:"JWT_EXPIRE_IN"`
	TokenAuth   *jwtauth.JWTAuth
}

func Initialize() {

	err := environment.Unmarshal(&CFG)
	if err != nil {
		panic(err)
	}
	os.Getenv("DB_DRIVER")
	CFG.Security.TokenAuth = jwtauth.New("HS256", []byte(CFG.Security.JWTSecret), nil)
}
