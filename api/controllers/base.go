package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// This import is needed to register postgres database - Don't remove the "underscore"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server : server struct (Gorm DB, Mux Router)
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize : Database connecton information, initialize our routes
func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v\n", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %v database\n", DbDriver)
			log.Fatal("This is the error : ", err)
		} else {
			fmt.Printf("We are connected to the %v database\n", DbDriver)
		}
	}

	server.Router = mux.NewRouter()

	server.InitializeRoutes()
}

// Run : Run the server on the specified port
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
