package logout

import (
	"fmt"

	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
)

type LogoutUC struct {
	Gateway LogoutGateway
}

func BuildLogoutUC(g LogoutGateway) *LogoutUC {
	return &LogoutUC{g}
}

func (bs *LogoutUC) Execute(encodedToken string) error {
	au, err := bs.Gateway.ExtractTokenMetadata(encodedToken)
	if err != nil {
		return fmt.Errorf(err_msg.UNAUTHORIZED)
	}

	deleted, err := bs.Gateway.DeleteAuth(au.AccessUUID)
	if err != nil || deleted == 0 {
		return fmt.Errorf(err_msg.UNAUTHORIZED)
	}

	return nil
}
