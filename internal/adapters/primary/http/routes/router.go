package routes

import (
	"github.com/axel-andrade/deu-role-auth/docs"
	"github.com/axel-andrade/deu-role-auth/internal/infra"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine, d *infra.Dependencies) *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Swapp API"
	docs.SwaggerInfo.Description = "This is a sample server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "swagg.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	main := r.Group("/")
	configureDefaultRoutes(main)

	v1 := r.Group("api/v1")
	configureAuthRoutes(v1, d)
	configureUsersRoutes(v1, d)

	return r
}
