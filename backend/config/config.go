package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppPort, DBHost, DBPort, DBUser, DBPassword, DBName, DBSSLMode, JWTSecret string
	JWTExpiryHours                                                            int
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found; using environment variables")
	}
	return &Config{
		AppPort: get("APP_PORT", "8081"), DBHost: get("DB_HOST", "localhost"), DBPort: get("DB_PORT", "5432"),
		DBUser: get("DB_USER", "postgres"), DBPassword: os.Getenv("DB_PASSWORD"), DBName: get("DB_NAME", "mh14"),
		DBSSLMode: get("DB_SSLMODE", "disable"), JWTSecret: get("JWT_SECRET", "change-me"),
		JWTExpiryHours: getInt("JWT_EXPIRY_HOURS", 24),
	}
}
func get(k, f string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return f
}
func getInt(k string, f int) int {
	v, e := strconv.Atoi(os.Getenv(k))
	if e != nil || v <= 0 {
		return f
	}
	return v
}
