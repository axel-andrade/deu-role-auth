package login

import (
	"fmt"
	"log"

	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
)

type LoginUC struct {
	Gateway LoginGateway
}

func BuildLoginUC(g LoginGateway) *LoginUC {
	return &LoginUC{g}
}

func (bs *LoginUC) Execute(input LoginInputDTO) (*LoginOutputDTO, error) {
	log.Println("Search already user with email: ", input.Email)
	user, err := bs.Gateway.FindUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf(err_msg.USER_NOT_FOUND)
	}

	log.Println("Comparing passwords")
	if err = bs.Gateway.CompareHashAndPassword(user.Password.Value, input.Password); err != nil {
		return nil, fmt.Errorf(err_msg.INVALID_PASSWORD)
	}

	log.Println("Generate token")
	td, err := bs.Gateway.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err = bs.Gateway.CreateAuth(user.ID, td); err != nil {
		return nil, err
	}

	return &LoginOutputDTO{td.AccessToken, td.RefreshToken}, nil
}
