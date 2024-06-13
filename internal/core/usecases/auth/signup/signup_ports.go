package signup

import "github.com/axel-andrade/deu-role-auth/internal/core/domain"

type SignupGateway interface {
	CancelTransaction() error
	CreateUser(user *domain.User) (*domain.User, error)
	CommitTransaction() error
	EncryptPassword(p string) (string, error)
	FindUserByEmail(email string) (*domain.User, error)
	StartTransaction() error
}

type SignupInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupOutputDTO struct {
	User domain.User
}
