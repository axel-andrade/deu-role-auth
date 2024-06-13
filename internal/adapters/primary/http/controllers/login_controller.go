package controllers

import (
	common_adapters "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/common"
	"github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/login"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUC   login.LoginUC
	Presenter presenters.LoginPresenter
}

func (ctrl *LoginController) Run(input login.LoginInputDTO) common_adapters.OutputPort {
	result, err := ctrl.LoginUC.Execute(input)
	return ctrl.Presenter.Show(result, err)
}

func BuildLoginController(uc *login.LoginUC, ptr *presenters.LoginPresenter) *LoginController {
	return &LoginController{LoginUC: *uc, Presenter: *ptr}
}

func (ctrl *LoginController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)
	loginInput := login.LoginInputDTO{
		Email:    inputMap["email"].(string),
		Password: inputMap["password"].(string),
	}

	result, err := ctrl.LoginUC.Execute(loginInput)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
