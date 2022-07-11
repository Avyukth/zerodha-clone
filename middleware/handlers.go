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
)


type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`

}

func CreateConnection() *sql.DB{
	err:= godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}


func CreateStock(w http.ResponseWriter, r *http.Request){
	var stock models.Stock
	err:=json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Error decoding request body. %v", err)
	}
	insertId := insertStock(stock)
	res := response{ID: insertId, Message: "Stock created successfully"}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) {
	panic("unimplemented")
}


func GetStock(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	id, err:=strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Error converting id to int. %v", err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatal("Unable to get Stock. %v", err)
	}


}

func GetAllStock(w http.ResponseWriter, r *http.Request){
	stocks, err := getAllStock()
	if err != nil {
		log.Fatal("Unable to get Stock. %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	id, err:=strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Error converting id to int. %v", err)
	}
	var stock models.Stock

	err:=json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Error decoding request body. %v", err)
	}
	updatedRows,err = updateStock(int64(id), stock)
	if err != nil {
		log.Fatal("Unable to update Stock. %v", err)
	}
	msg:= fmt.Sprintf("Stock Updated successfully. Total rows/records affected: %v", updatedRows)

	res := response{ID: int64(id), Message: msg}
	json.NewEncoder(w).Encode(res)
}

func updateStock(i int64, stock models.Stock) {
	panic("unimplemented")
}

func DeleteStock(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	id, err:=strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Error converting id to int. %v", err)
	}
	deletedRows,err := deleteStock(int64(id))
	if err != nil {
		log.Fatal("Unable to delete Stock. %v", err)
	}
	msg:= fmt.Sprintf("Stock Deleted successfully. Total rows/records affected: %v", deletedRows)
	res := response{ID: int64(id), Message: msg}
	json.NewEncoder(w).Encode(res)
}
