package login

import "github.com/axel-andrade/deu-role-auth/internal/core/domain"

type LoginGateway interface {
	CreateAuth(userid string, td *domain.TokenDetails) error
	CompareHashAndPassword(hash string, p string) error
	FindUserByEmail(email string) (*domain.User, error)
	GenerateToken(userid string) (*domain.TokenDetails, error)
}

type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
