package router

import (
	"github.com/gorilla/mux"
	"github.com/ishanshre/GO-Stocks-API/pkg/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stocks/all", middleware.GetStocks).Methods("GET")
	router.HandleFunc("/api/stocks/{id}", middleware.GetStockByID).Methods("GET")
	router.HandleFunc("/api/stocks/create", middleware.CreateStock).Methods("POST")
	router.HandleFunc("/api/stocks/{id}/update", middleware.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stocks/{id}/delete", middleware.DeleteStock).Methods("DELETE")
	return router
}
