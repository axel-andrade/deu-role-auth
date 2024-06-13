package controllers

import (
	"github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/presenters"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/signup"
	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	SignupUC  signup.SignupUC
	Presenter presenters.SignupPresenter
}

func BuildSignUpController(uc *signup.SignupUC, ptr *presenters.SignupPresenter) *SignUpController {
	return &SignUpController{SignupUC: *uc, Presenter: *ptr}
}

// @Summary		Register user
// @Description	Register an user.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			body	body		signup.SignupInputDTO	true	"Corpo da solicitação"
// @Success		200		{object}	common_ptr.UserFormatted
// @Failure		400		{object}	shared_err.InvalidOperationError	"Bad Request"
// @Failure		500		{object}	shared_err.InternalError			"Internal Server Error"
// @Router			/api/v1/auth/signup [post]
func (ctrl *SignUpController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)
	signupInput := signup.SignupInputDTO{
		Email:    inputMap["email"].(string),
		Name:     inputMap["name"].(string),
		Password: inputMap["password"].(string),
	}

	result, err := ctrl.SignupUC.Execute(signupInput)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
