package models

import "time"

type Transaction struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Description string    `json:"description,omitempty"`
	Color       string    `json:"color,omitempty"`
	User        User      `json:"user"`
	Date        time.Time `json:"date"`
	Account     Account   `json:"account"`
	Category    Category  `json:"category"`
	UserID      int
	AccountID   int
	CategoryID  int
}

func (Transaction) TableName() string {
	return "transaction"
}
