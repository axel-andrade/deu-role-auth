package presenters

import (
	"net/http"

	cad "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/common"

	common_ptr "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters/common"
	"github.com/axel-andrade/deu-role-auth/internal/core/domain"
	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/get_users"
)

type GetUsersOutputFormatted struct {
	Users      []common_ptr.UserFormatted `json:"users"`
	Pagination common_ptr.PaginateResult  `json:"pagination"`
}

type GetUsersPresenter struct {
	userPtr       common_ptr.UserPresenter
	paginationPtr common_ptr.PaginationPresenter
}

func BuildGetUsersPresenter() *GetUsersPresenter {
	return &GetUsersPresenter{}
}

func (ptr *GetUsersPresenter) Show(result *get_users.GetUsersOutputDTO, paginationOptions domain.PaginationOptions, err error) cad.OutputPort {
	if err != nil {
		return cad.OutputPort{StatusCode: http.StatusBadRequest, Data: cad.ErrorMessage{Message: err_msg.INTERNAL_SERVER_ERROR}}
	}

	u := ptr.userPtr.FormatList(result.Users)
	p := ptr.paginationPtr.Format(paginationOptions, result.TotalUsers)
	data := GetUsersOutputFormatted{Users: u, Pagination: p}

	return cad.OutputPort{StatusCode: http.StatusOK, Data: data}
}
