package logout

import "github.com/axel-andrade/deu-role-auth/internal/core/domain"

type LogoutGateway interface {
	ExtractTokenMetadata(encoded string) (*domain.AccessDetails, error)
	DeleteAuth(uuid string) (int64, error)
}
