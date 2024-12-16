package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	ENV      string         `yaml:"env" env-required:"true"`
	Server   ServerConfig   `yaml:"server" env-required:"true"`
	Redis    RedisConfig    `yaml:"redis" env-required:"true"`
	Postgres PostgresConfig `yaml:"postgres" env-required:"true"`
	Auth     AuthConfig     `yaml:"auth" env-required:"true"`
	Crypto   CryptoConfig   `yaml:"crypto" env-required:"true"`
	SMTP     SMTPConfig     `yaml:"smtp" env-required:"true"`
}

type ServerConfig struct {
	Port string `yaml:"port" env-default:"8080"`
}

type RedisConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db" env-default:"0"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"db_name" env-required:"true"`
	SSLMode  bool   `yaml:"ssl_mode" env-default:"false"`
}

type AuthConfig struct {
	AccessJWTTTL  string `yaml:"access_jwt_ttl" env-required:"true"`
	RefreshJWTTTL string `yaml:"refresh_jwt_ttl" env-required:"true"`
	JWTSecretKey  string `yaml:"jwt_secret_key" env-required:"true"`
}

type CryptoConfig struct {
	PasswordKey        string `yaml:"password_key" env-required:"true"`
	AuthCredentialsKey string `yaml:"auth_credentials_key" env-required:"true"`
}

type SMTPConfig struct {
	Host          string `yaml:"host" env-required:"true"`
	Port          string `yaml:"port" env-required:"true"`
	FromEmail     string `yaml:"from_email" env-required:"true"`
	FromEmailName string `yaml:"from_email_name" env-required:"true"`
}

// MustLoad is going to panic if loader return error
func MustLoad(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file %s does not exist", path))
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
