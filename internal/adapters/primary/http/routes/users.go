package routes

import (
	"github.com/axel-andrade/deu-role-auth/internal/adapters/primary/http/middlewares"
	"github.com/axel-andrade/deu-role-auth/internal/infra"
	"github.com/gin-gonic/gin"
)

func configureUsersRoutes(r *gin.RouterGroup, d *infra.Dependencies) {
	users := r.Group("users")
	{
		users.GET("/", middlewares.Authorize(d), middlewares.ValidateRequest("users/get_users"), d.GetUsersController.Handle)
	}
}
