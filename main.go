package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heydp/oms/pkg/db"
	"github.com/heydp/oms/pkg/handlers"
)

var portnumber = ":8080"

func main() {
	dbCmd := flag.NewFlagSet("db", flag.ExitOnError)
	DB := db.Init(dbCmd)
	h := handlers.NewHandler(DB)

	r := mux.NewRouter()
	r.HandleFunc("/orders", h.GetOrdersHandler).Methods(http.MethodGet)
	r.HandleFunc("/order/{id}", h.GetOrderHandler).Methods(http.MethodGet)

	r.HandleFunc("/orders", h.PostOrderHandler).Methods(http.MethodPost)
	r.HandleFunc("/order/{id}", h.UpdateOrderHandler).Methods(http.MethodPatch)

	// r.HandleFunc("/order/{id}", h.DeleteOrderHandler).Methods(http.MethodDelete)

	log.Println("Server Up and Running at PortNumber ", portnumber)
	err := http.ListenAndServe(portnumber, r)
	log.Fatal(err)

}
