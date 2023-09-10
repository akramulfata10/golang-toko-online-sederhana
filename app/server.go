package app

import (
	"flag"
	"log"
	"os"
	"flag"

	"github.com/akramulfata10/gotoko/app/controllers"
	"github.com/joho/godotenv"
<<<<<<< HEAD
)

=======
	"github.com/urfave/cli"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"github.com/akramulfata10/gotoko/database/seeders"

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

func (server *Server) Initialize(appConfig AppConfig, dbConfig DbConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)

	server.initializeRoutes()
}

func (server *Server) initializeDB(dbConfig DbConfig){

	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}

} 

func (server *Server) dbMigrate() {
	for _, model := range RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully.")
}


func (server *Server) initCommands(config AppConfig, dbConfig DbConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}



func (server *Server) Run(address string) {
	fmt.Printf("listening to port  %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

>>>>>>> d7c111a112f70f719ca912edeb2d438d4d53ae98
func getEnv(key, fallback string) string {
	if value, berhasil := os.LookupEnv(key); berhasil {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "gotoko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9191")
	appConfig.AppUrl = getEnv("APP_URL", "http://localhost:9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "gotoko-db")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
<<<<<<< HEAD
		server.InitCommands(appConfig, dbConfig)
=======
		server.initCommands(appConfig, dbConfig)
>>>>>>> d7c111a112f70f719ca912edeb2d438d4d53ae98
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

	// server.Initialize(appConfig, dbConfig)
	// server.Run(":" + appConfig.AppPort)
}
