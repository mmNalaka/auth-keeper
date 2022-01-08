package config

import "github.com/ilyakaznacheev/cleanenv"

type App struct {
	Name          string `env-required:"true" env:"APP_NAME"`
	Version       string `env-required:"true" env:"APP_VERSION"`
	DisableLogin  bool   `env:"DISABLE_LOGIN" env-default:"false"`
	DisableSignup bool   `env:"DISABLE_SIGNUP" env-default:"false"`
}

type Server struct {
	Port    string `env-required:"true" env:"SERVER_PORT"`
	Timeout int    `env-required:"true" env:"SERVER_TIMEOUT_SECONDS"`
}

type Postgres struct {
	PoolMax       int    `env-required:"true" env:"PG_POOL_MAX"`
	URL           string `env-required:"true" env:"PG_URL"`
	RunMigrations bool   `env:"PG_RUN_MIGRATIONS" env-default:"false"`
}

type Redis struct {
	Host     string `env-required:"true" env:"REDIS_HOST"`
	Port     int    `env-required:"true" env:"REDIS_PORT"`
	Password string `env-required:"true" env:"REDIS_PASSWORD"`
}

type JWT struct {
	Secret    string `env-required:"true" env:"JWT_SECRET"`
	Algorithm string `env-required:"true" env:"JWT_ALGORITHM"`
	Expires   int    `env-required:"true" env:"JWT_EXPIRES_IN_SECONDS"`
}

type Log struct {
	Level string `env-required:"true" env:"LOG_LEVEL"`
}

type Config struct {
	App      App
	Server   Server
	Postgres Postgres
	Redis    Redis
	JWT      JWT
	Log      Log
}

var Cfg Config

func Init() {
	if err := cleanenv.ReadEnv(&Cfg); err != nil {
		panic(err)
	}
}
