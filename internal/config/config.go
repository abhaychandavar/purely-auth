package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	port                = os.Getenv("PORT")
	dbHost              = os.Getenv("MONGO_DB_HOST")
	dbPort              = os.Getenv("MONGO_DB_PORT")
	db                  = os.Getenv("MONGO_DB_DATABASE")
	internalAccessToken = os.Getenv("INTERNAL_ACCESS_TOKEN")
	firebaseConfigPath  = os.Getenv("FIREBASE_CONFIG_PATH")
)

type configType struct {
	Port                string
	DbHost              string
	DbPort              string
	Db                  string
	InternalAccessToken string
	FirebaseConfigPath  string
}

func GetConfig() configType {
	obj := configType{
		DbHost:              dbHost,
		DbPort:              dbPort,
		Db:                  db,
		Port:                "8080",
		InternalAccessToken: internalAccessToken,
		FirebaseConfigPath:  firebaseConfigPath,
	}
	if port != "" {
		obj.Port = port
	}
	return obj
}
