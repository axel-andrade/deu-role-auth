package get_users

import "github.com/axel-andrade/deu-role-auth/internal/core/domain"

type GetUsersGateway interface {
	GetUsersPaginate(pagination domain.PaginationOptions) ([]domain.User, uint64, error)
}

type GetUsersInputDTO struct {
	PaginationOptions domain.PaginationOptions
}

type GetUsersOutputDTO struct {
	Users      []domain.User
	TotalUsers uint64
}
