package config

import (
	"os"
)

func GetDSN() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user:userpassword@tcp(localhost:3306)/library_db?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return databaseURL
}
