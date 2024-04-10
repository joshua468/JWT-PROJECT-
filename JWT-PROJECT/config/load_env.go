package config

import "time"

type Config struct {
	DBHOST     string `mapstructure:"POSTGRES_HOST"`
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_DB"`
	DBPORT     string `mapstructure:"POSTGRES_PORT"`

	Serverport string `mapstructure:"PORT"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_ExpiresIn"`
	TokenMaxAge    string        `mapstructure:""`
}
