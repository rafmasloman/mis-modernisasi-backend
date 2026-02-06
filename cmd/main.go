package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/config"
	"github.com/rafmasloman/mis-modernisasi-backend/internal/core"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error .env failed to load")
	}

	fmt.Println("Hello World")

	db, err := config.GormApp()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	v1 := r.Group("/api/v1/mis-new")

	core.RegisterReportBuilderCore(v1, db)
	core.RegisterFilterBuilderCore(v1, db)

	AppHost := os.Getenv("APP_HOST")
	AppPort := os.Getenv("APP_PORT")
	AppRun := fmt.Sprintf("%s:%s", AppHost, AppPort)

	r.Run(AppRun)
}
