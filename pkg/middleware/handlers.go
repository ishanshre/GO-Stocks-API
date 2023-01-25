package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/ishanshre/GO-Stocks-API/pkg/models"
	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error Loading environment files")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successfull")
	return db
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()
	if err != nil {
		log.Fatalf("Error retriving all the stocks, %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func GetStockByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Error converting string to int", err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("unable to get stock, %v", err)
	}
	json.NewEncoder(w).Encode(stock)
}
func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request-> %v", err)
	}
	insertId := insertStock(stock)
	res := response{
		ID:      insertId,
		Message: "Stock created Successfully",
	}
	json.NewEncoder(w).Encode(res)

}
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Error in converting string to interger, %v", err)
	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Error in parsing the request from the client, %v", err)
	}
	updatedRows := updateStock(int64(stockId), stock)
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected : %v", updatedRows)
	res := response{
		ID:      int64(stockId),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)

}
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Error in converting string to int, %v", err)
	}
	deleteRows := deleteStock(int64(stockId))
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/record affected %v", deleteRows)
	res := response{
		ID:      int64(stockId),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
