package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Active   bool   `json:"active,omitempty"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles"`
}

func (User) TableName() string {
	return "user"
}
