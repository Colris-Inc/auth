package config

import (
	"log"
	"os"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	EncryptionKey string
}

var AppConfig *Config

func StartupConfiguration() {
	log.Println("Configuring appliaction")
	log.Println("Reading and setting env variables")
	AppConfig = &Config{
		DBHost:        os.Getenv("DBHOST"),
		DBPort:        os.Getenv("DBPORT"),
		DBUser:        os.Getenv("DBUSER"),
		DBPassword:    os.Getenv("DBPASSWORD"),
		DBName:        os.Getenv("DBNAME"),
		EncryptionKey: os.Getenv("ENCRYPTIONKEY"),
	}
	log.Panicln("configuration complete")
}
