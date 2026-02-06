package config

import (
	"os"
)

type EnvGormType struct {
	DBHost string
	DBUser string
	DBName string
	DBPort string
	DBPass string
}

type EnvGeneralType struct {
	BLUrl string
}

func EnvGorm() *EnvGormType {

	dbEnv := EnvGormType{
		DBHost: os.Getenv("DBHOST"),
		DBUser: os.Getenv("DBUSER"),
		DBName: os.Getenv("DBNAME"),
		DBPort: os.Getenv("DBPORT"),
		DBPass: os.Getenv("DBPASS"),
	}

	return &dbEnv
}
