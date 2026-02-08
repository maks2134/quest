package configs

import (
	"errors"
	"io/fs"

	"github.com/ilyakaznacheev/cleanenv"
)

var Configs config

type config struct {
	DBHost           string `env:"DB_HOST" env-default:"localhost"`
	DBPort           string `env:"DB_PORT" env-default:"5432"`
	DBUser           string `env:"POSTGRES_USER" env-default:"quest"`
	DBPassword       string `env:"POSTGRES_PASSWORD" env-default:"quest"`
	DBName           string `env:"POSTGRES_DB" env-default:"quest"`
	DBSSLMode        string `env:"DB_SSLMODE" env-default:"disable"`
	CORSAllowOrigins string `env:"CORS_ALLOW_ORIGINS" env-default:"http://localhost:5173,http://localhost:3000"`
	CORSAllowMethods string `env:"CORS_ALLOW_METHODS" env-default:"GET,POST,PUT,DELETE,OPTIONS"`
	CORSAllowHeaders string `env:"CORS_ALLOW_HEADERS" env-default:"Origin,Content-Type,Accept,Authorization"`
	CORSMaxAge       int    `env:"CORS_MAX_AGE" env-default:"3600"`
	SwaggerUser      string `env:"SWAGGER_USER" env-default:"admin"`
	SwaggerPassword  string `env:"SWAGGER_PASSWORD" env-default:"admin"`
}

func LoadConfig() {
	if err := cleanenv.ReadConfig(".env", &Configs); err != nil && !errors.Is(err, fs.ErrNotExist) {
		panic(err)
	}
	if err := cleanenv.ReadEnv(&Configs); err != nil {
		panic(err)
	}
}
