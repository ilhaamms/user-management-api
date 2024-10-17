package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilhaamms/user-management-api/config"
)

func main() {

	mux := http.NewServeMux()

	db, err := config.GetConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.Close()

	mux.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Println("Server started on :8080")

	http.ListenAndServe(":8080", mux)

}
