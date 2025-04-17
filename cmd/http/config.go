package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBConnURL  string
}

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := Config{}
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.DBConnURL = os.Getenv("DB_CONN_URL")
	return config
}
