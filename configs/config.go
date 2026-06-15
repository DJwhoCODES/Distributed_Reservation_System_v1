package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	RedisAddr string

	HoldTTLSeconds int
}

func Load() *Config {
	_ = godotenv.Load()

	ttl, err := strconv.Atoi(getEnv("HOLD_TTL_SECONDS", "120"))
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		AppPort: getEnv("APP_PORT", "8000"),

		PostgresHost:     getEnv("POSTGRES_HOST", "127.0.0.1"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:       getEnv("POSTGRES_DB", "ticket-reservation"),

		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),

		HoldTTLSeconds: ttl,
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
