package models

import "time"

// Transaction структура для хранения информации о транзакциях
type Transaction struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
