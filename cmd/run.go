package cmd

import (
	"exam/handlers"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateTransactionHandler(w, r)
		case http.MethodGet:
			handlers.GetAllTransactionsHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/transactions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTransactionByIDHandler(w, r)
		case http.MethodPut:
			handlers.UpdateTransactionByIDHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteTransactionByIDHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
