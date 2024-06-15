package infra

import (
	grpc_services "github.com/axel-andrade/deu-role-auth/internal/adapters/primary/grpc/services"
	"github.com/axel-andrade/deu-role-auth/internal/adapters/primary/handlers"
	mongo_repositories "github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo/repositories"
	redis_repositories "github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/redis/repositories"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/login"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/logout"
	"github.com/axel-andrade/deu-role-auth/internal/core/usecases/auth/signup"
)

type Dependencies struct {
	UserRepository    *mongo_repositories.UserRepository
	SessionRepository *redis_repositories.SessionRepository

	EncrypterHandler    *handlers.EncrypterHandler
	TokenManagerHandler *handlers.TokenManagerHandler

	SignupUC *signup.SignupUC
	LoginUC  *login.LoginUC
	LogoutUC *logout.LogoutUC

	AuthGrpcService *grpc_services.AuthGrpcService
}

func LoadDependencies() *Dependencies {
	d := &Dependencies{}

	loadRepositories(d)
	loadHandlers(d)
	loadUseCases(d)
	loadGrpcServices(d)

	return d
}

func loadRepositories(d *Dependencies) {
	d.UserRepository = mongo_repositories.BuildUserRepository()
	d.SessionRepository = redis_repositories.BuildSessionRepository()
}

func loadHandlers(d *Dependencies) {
	d.EncrypterHandler = handlers.BuildEncrypterHandler()
	d.TokenManagerHandler = handlers.BuildTokenManagerHandler()
}

func loadUseCases(d *Dependencies) {
	d.SignupUC = signup.BuildSignupUC(struct {
		*mongo_repositories.UserRepository
		*handlers.EncrypterHandler
	}{d.UserRepository, d.EncrypterHandler})

	d.LoginUC = login.BuildLoginUC(struct {
		*redis_repositories.SessionRepository
		*mongo_repositories.UserRepository
		*handlers.EncrypterHandler
		*handlers.TokenManagerHandler
	}{d.SessionRepository, d.UserRepository, d.EncrypterHandler, d.TokenManagerHandler})

	d.LogoutUC = logout.BuildLogoutUC(struct {
		*redis_repositories.SessionRepository
		*handlers.TokenManagerHandler
	}{d.SessionRepository, d.TokenManagerHandler})
}

func loadGrpcServices(d *Dependencies) {
	d.AuthGrpcService = grpc_services.NewAuthGrpcService(*d.LoginUC, *d.SignupUC)
}
