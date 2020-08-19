package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type app struct {
	Name string
	Key  string
}

var App = &app{}

func (c *app) Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	App.Name = os.Getenv("APP_NAME")
	App.Key  = os.Getenv("APP_KEY")
	return nil
}
