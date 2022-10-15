package httpserver

import (
	"Assignment2/httpserver/controllers"
	"github.com/gorilla/mux"

)

func Router() *mux.Router  {
	router := mux.NewRouter()

	router.HandleFunc("/order", controllers.GetTheOrder ).Methods("GET")
	router.HandleFunc("/order/{Order_ID}", controllers.GetOrder ).Methods("GET")
	router.HandleFunc("/order", controllers.CreateOrder ).Methods("POST")
	router.HandleFunc("/order/{Order_ID}", controllers.UpdateOrder ).Methods("POST")
	router.HandleFunc("/order/{Order_ID}", controllers.DeleteOrder ).Methods("DELETE")
	return router
}