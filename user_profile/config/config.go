package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"

)

type Config struct {
	CollaborationServicePort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.CollaborationServicePort = cast.ToString(coalesce("COLLABORATION_SERVICE_PORT", ":8083"))

	config.PostgresHost = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(coalesce("DB_PORT", 5432))
	config.PostgresUser = cast.ToString(coalesce("DB_USER", "postgres"))
	config.PostgresPassword = cast.ToString(coalesce("DB_PASSWORD", "0509"))
	config.PostgresDatabase = cast.ToString(coalesce("DB_NAME", "reja"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}

// CREATE TYPE user_role AS ENUM ('musician', 'listener', 'producer');

// CREATE TABLE user_profiles (
//     user_id INTEGER PRIMARY KEY REFERENCES users(id),
//     full_name VARCHAR(100),
//     bio TEXT,
//     role user_role,
//     location VARCHAR(100),
//     avatar_url VARCHAR(255),
//     website VARCHAR(255)
// );
