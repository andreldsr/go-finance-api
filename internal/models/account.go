package models

type Account struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name,omitempty"`
	Color  string `json:"color,omitempty"`
	UserID int
	User   User `json:"user"`
}

func (Account) TableName() string {
	return "account"
}
