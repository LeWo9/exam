package handlers

import (
	"encoding/json"
	"exam/models"
	"net/http"
	"strconv"
	"time"
)

var transactions []models.Transaction
var idCounter = 1

// Handlers
func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var newTransaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTransaction.ID = idCounter
	idCounter++
	newTransaction.Date = time.Now()

	transactions = append(transactions, newTransaction)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTransaction)
}

func GetAllTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/transactions/"):])

	for _, transaction := range transactions {
		if transaction.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transaction)
			return
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}

func UpdateTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/transactions/"):])

	var updatedTransaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&updatedTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, transaction := range transactions {
		if transaction.ID == id {
			updatedTransaction.ID = id
			updatedTransaction.Date = time.Now()
			transactions[i] = updatedTransaction

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTransaction)
			return
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}

func DeleteTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/transactions/"):])

	for i, transaction := range transactions {
		if transaction.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Transaction not found", http.StatusNotFound)
}
