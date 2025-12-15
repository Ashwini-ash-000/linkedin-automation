package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Email               string
	Password            string
	Headless            bool
	MaxDailyConnections int
	WorkStartHour       int
	WorkEndHour         int
}

func Load() *Config {
	cfg := &Config{
		Email:               os.Getenv("BOT_EMAIL"),
		Password:            os.Getenv("BOT_PASSWORD"),
		Headless:            getBool("HEADLESS", false),
		MaxDailyConnections: getInt("MAX_DAILY_CONNECTIONS", 15),
		WorkStartHour:       getInt("WORK_START_HOUR", 10),
		WorkEndHour:         getInt("WORK_END_HOUR", 18),
	}

	if cfg.Email == "" || cfg.Password == "" {
		log.Println("Warning: BOT_EMAIL or BOT_PASSWORD not set (expected for mock/demo)")
	}

	return cfg
}

func getInt(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return def
}

func getBool(key string, def bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return def
}
