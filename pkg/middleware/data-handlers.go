package middleware

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ishanshre/GO-Stocks-API/pkg/models"
)

func getAllStock() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stocks []models.Stock
	sqlStatement := "SELECT * FROM stocks"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the 	query")
	}
	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}

func getStock(id int64) (models.Stock, error) {
	var stock models.Stock
	db := createConnection()
	defer db.Close()
	sqlStatement := "SELECT * FROM stocks WHERE stockID=$1"
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row %v", err)
	}
	return stock, err
}

func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatememnt := "INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockID"
	var id int64
	// here $1 maps stock.Name, $2 maps stock.Price, $3 maps stock.Company
	err := db.QueryRow(sqlStatememnt, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("unable to execute the query %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}

func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := "UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockID=$1"
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to check the rows affected, %v", err)
	}
	fmt.Printf("Total rows/records affected, %v", rowsAffected)
	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := "DELETE FROM stocks WHERE stockID=$1"
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query, %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to read the rowsAffected, %v", err)
	}
	fmt.Printf("No. of rows/recored affeted, %v", rowsAffected)
	return rowsAffected
}
