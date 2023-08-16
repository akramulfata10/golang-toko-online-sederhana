package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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

type DbConfig struct {
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
	DBDriver string
}

func (server *Server) KoneksiDB(appConfig AppConfig, dbConfig DbConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)

	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Could not connect to database")
	}
	
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("listening to port  %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func getEnv(key, fallback string) string {
	if value, berhasil := os.LookupEnv(key); berhasil {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DbConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}

	appConfig.AppName = getEnv("APP_NAME","gotoko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT","9191")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "gotoko-db")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	server.KoneksiDB(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
