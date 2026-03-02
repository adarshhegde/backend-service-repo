package config

import "github.com/caarlos0/env/v11"

type Config struct {
	MongoURI string `env:"MONGODB_URI,required"`
	HTTPPort int    `env:"HTTP_PORT" envDefault:"8080"`
	GRPCPort int    `env:"GRPC_PORT" envDefault:"3001"`
}

func Load() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return cfg, err
}
