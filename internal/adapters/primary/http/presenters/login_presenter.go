package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/login"
)

type LoginOutputFormatted struct {
	AccessToken  string                   `json:"access_token" description:"Access token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	RefreshToken string                   `json:"refresh" description:"Access token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	User         common_ptr.UserFormatted `json:"user"`
}

type LoginPresenter struct {
	userPtr common_ptr.UserPresenter
}

func BuildLoginPresenter() *LoginPresenter {
	return &LoginPresenter{}
}

func (p *LoginPresenter) Show(result *login.LoginOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return p.formatSuccessOutput(result)
}

func (p *LoginPresenter) formatSuccessOutput(result *login.LoginOutputDTO) common_adapters.OutputPort {
	data := LoginOutputFormatted{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		User:         p.userPtr.Format(result.User),
	}

	return common_adapters.OutputPort{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func (p *LoginPresenter) formatError(err error) common_adapters.OutputPort {
	switch err.Error() {
	case err_msg.USER_NOT_FOUND:
		return common_adapters.OutputPort{StatusCode: http.StatusNotFound, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	case err_msg.INVALID_PASSWORD:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	default:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
	}
}
