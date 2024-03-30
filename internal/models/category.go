package models

type Category struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name,omitempty"`
	Color     string `json:"color,omitempty"`
}

func (Category) TableName() string {
	return "category"
}
