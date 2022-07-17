package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"zerodha-clone/routers"
)

func main() {
	var port string
	if os.Getenv("ENV") == "PROD" {
		port = os.Getenv("PROD_APP_PORT")

	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		port = os.Getenv("DEV_APP_PORT")
	}
	portString := ":" + port
	r := routers.Router()
	fmt.Printf("Starting server on port %s..........\n", port)
	log.Fatal(http.ListenAndServe(portString, r))
}
