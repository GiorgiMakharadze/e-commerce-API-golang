package validators

import (
	"errors"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
)

func ValidateUserRole(role string) (models.UserRole, error) {
	switch role {
	case string(models.Admin), string(models.Seller), string(models.Buyer):
		return models.UserRole(role), nil
	default:
		return "", errors.New("invalid user role")
	}
}
