package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv string
	AppPort string
}

func (server *Server) KoneksiDB(appConfig AppConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("listening to port  %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}

	appConfig.AppName = getEnv("APP_NAME","gotoko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT","9191")

	server.KoneksiDB(appConfig)
	server.Run(":" + appConfig.AppPort)
}
