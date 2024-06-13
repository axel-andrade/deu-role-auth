package presenters

import (
	"net/http"

	common_adapters "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/common"
	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
)

type LogoutPresenter struct{}

func BuildLogoutPresenter() *LogoutPresenter {
	return &LogoutPresenter{}
}

func (p *LogoutPresenter) Show(err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return common_adapters.OutputPort{StatusCode: http.StatusNoContent}
}

func (p *LogoutPresenter) formatError(err error) common_adapters.OutputPort {
	switch err.Error() {
	case err_msg.UNAUTHORIZED:
		return common_adapters.OutputPort{StatusCode: http.StatusUnauthorized, Data: common_adapters.ErrorMessage{Message: err.Error()}}
	default:
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
	}
}
