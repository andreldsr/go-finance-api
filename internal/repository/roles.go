package repository

import (
	"go-finance-api/internal/database"
	"go-finance-api/internal/models"
)

func FindRolesByName(names []string) (roles []models.Role) {
	database.DB.Model(&models.Role{}).Where("name in ?", names).Scan(&roles)
	return
}
