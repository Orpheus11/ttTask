package main

import (
	"log"
	"net/http"

	"github.com/goconf/conf"
	"github.com/gorilla/mux"
	"gopkg.in/pg.v3"
	"ttTask/controller"
	"ttTask/model"
)

var (
	config *conf.ConfigFile
	listen string
	DB     *pg.DB
)

func init() {
	config, _ := conf.ReadConfigFile("config.conf")
	listen, _ = config.GetString("default", "listen")
	listen = ":" + listen
	db_user, _ := config.GetString("default", "db_user")
	db_passwd, _ := config.GetString("default", "db_passwd")
	model.InitConnect(db_user, db_passwd)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index)
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users", controller.AddUsers).Methods("POST")
	router.HandleFunc("/users/{user_id}/relationships", controller.GetRelationshipsByUserId).Methods("GET")
	router.HandleFunc("/users/{user_id}/relationships/{other_user_id}", controller.AddRelationships).Methods("PUT")
	log.Fatal(http.ListenAndServe(listen, router))
}
