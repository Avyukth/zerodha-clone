package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"zerodha-clone/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_DSN_URL"))
	if err != nil {
		log.Fatal("Error getting db connection ............")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error getting db connection ............in ping")
		panic(err)
	}
	fmt.Println("Successfully connected!............")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Error decoding request body. %v", err)
	}
	insertId := insertStock(stock)
	res := response{ID: insertId, Message: "Stock created successfully"}
	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error converting id to int. %v", err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get Stock. %v", err)
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()
	if err != nil {
		log.Fatalf("Unable to get Stock. %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error converting id to int. %v", err)
	}
	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Error decoding request body. %v", err)
	}
	updatedRows := updateStock(int64(id), stock)
	if err != nil {
		log.Fatalf("Unable to update Stock. %v", err)
	}
	msg := fmt.Sprintf("Stock Updated successfully. Total rows/records affected: %v", updatedRows)

	res := response{ID: int64(id), Message: msg}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error converting id to int. %v", err)
	}
	deletedRows := deleteStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to delete Stock. %v", err)
	}
	msg := fmt.Sprintf("Stock Deleted successfully. Total rows/records affected: %v", deletedRows)
	res := response{ID: int64(id), Message: msg}
	json.NewEncoder(w).Encode(res)
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()

	defer db.Close()
	sqlStatement := `SELECT * FROM stocks WHERE stock_id=$1`
	var stock models.Stock
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return stock, nil
}

func insertStock(stock models.Stock) int64 {
	db := CreateConnection()

	defer db.Close()
	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stock_id`
	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Error in executing the Query. %v", err)
	}
	return id
}

func getAllStock() ([]models.Stock, error) {
	db := CreateConnection()

	defer db.Close()
	sqlStatement := `SELECT * FROM stocks`
	var stocks []models.Stock
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Error in executing the Query. %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

func updateStock(id int64, stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE stocks SET name=$1, price=$2, company=$3 WHERE stock_id=$4 RETURNING stock_id`
	res, err := db.Exec(sqlStatement, stock.Name, stock.Price, stock.Company, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to get the rows affected. %v", err)
	}
	fmt.Printf("Total rows/records affected: %v", rowsAffected)
	return rowsAffected

}

func deleteStock(id int64) int64 {
	db := CreateConnection()

	defer db.Close()
	sqlStatement := `DELETE FROM stocks WHERE stock_id=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to get the rows affected. %v", err)
	}
	fmt.Printf("Total rows/records affected: %v", rowsAffected)
	return rowsAffected
}
