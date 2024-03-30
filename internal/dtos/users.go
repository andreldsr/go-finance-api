package dtos

type UserListDto struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserDto struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
