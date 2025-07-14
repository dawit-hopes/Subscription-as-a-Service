package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	MongoDBConnectionString string `env:"MONGODB_CONNECTION_STRING,required"`
	MongoDBName             string `env:"MONGODB_NAME,required"`
	SigningKey              string `env:"SIGNING_KEY,required"`
	Issuer                  string `env:"ISSUER,required"`
	Cost                    int    `env:"COST" envDefault:"10"`
	ServerAddress           string `env:"SERVER_ADDRESS" envDefault:":8080"`
}

func NewEnv() *Env {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables.")
	}

	cfg := &Env{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v", err)
	}

	return cfg
}
