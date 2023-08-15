package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) KoneksiDB() {
	fmt.Println("Welcome To Gotoko")
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("listening to port  %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func Run() {
	var testKoneksi = Server{}
	testKoneksi.KoneksiDB()
	testKoneksi.Run(":9000")
}
