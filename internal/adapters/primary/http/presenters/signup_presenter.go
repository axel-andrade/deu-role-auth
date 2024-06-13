package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/common"
	common_ptr "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters/common"
	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/signup"
)

type SignupPresenter struct {
	UserPtr common_ptr.UserPresenter
}

func BuildSignupPresenter() *SignupPresenter {
	return &SignupPresenter{}
}

func (p *SignupPresenter) Show(result *signup.SignupOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return common_adapters.OutputPort{StatusCode: http.StatusCreated, Data: p.UserPtr.Format(result.User)}
}

func (p *SignupPresenter) formatError(err error) common_adapters.OutputPort {
	switch err.Error() {
	case err_msg.INVALID_PASSWORD, err_msg.INVALID_EMAIL, err_msg.INVALID_NAME:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	case err_msg.USER_ALREADY_EXISTS:
		return common_adapters.OutputPort{StatusCode: http.StatusConflict, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	default:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
	}
}
