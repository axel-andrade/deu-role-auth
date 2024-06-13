package controllers

import (
	"github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/logout"
	interactor "github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/logout"
	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUC  interactor.LogoutUC
	Presenter presenters.LogoutPresenter
}

func BuildLogoutController(uc *logout.LogoutUC, logoutPtr *presenters.LogoutPresenter) *LogoutController {
	return &LogoutController{LogoutUC: *uc, Presenter: *logoutPtr}
}

func (ctrl *LogoutController) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	encodedToken := authHeader[len("Bearer "):]

	err := ctrl.LogoutUC.Execute(encodedToken)
	output := ctrl.Presenter.Show(err)

	c.JSON(output.StatusCode, output.Data)
}
