package mongo_repositories

import (
	"github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo/mappers"
	"github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo/models"
	"github.com/axel-andrade/deu-role-auth/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Base       *BaseRepository
	userMapper mappers.UserMapper
}

const collection = "users"

func BuildUserRepository() *UserRepository {
	userMapper := mappers.BuildUserMapper(&mappers.BaseMapper{})
	baseRepo := BuildBaseRepository(collection)

	return &UserRepository{baseRepo, *userMapper}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	model := r.userMapper.ToPersistence(*user)
	_, err := r.Base.Create(model)

	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomain(model), nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	// err := r.Db.Save(user).Error
	// return err
	return nil
}

func (r *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	filter := bson.M{"email": email}

	var model models.UserModel
	if err := r.Base.FindOne(filter).Decode(&model); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return r.userMapper.ToDomain(model), nil
}

func (r *UserRepository) FindUserByID(id string) (*domain.User, error) {
	// var user entities.User
	// err := r.Db.First(&user, "id = ?", id).Error

	// if err != nil || user.ID == "" {
	// 	return nil, err
	// }

	// return &user, nil
	return nil, nil
}

func (r *UserRepository) GetUsersPaginate(pagination domain.PaginationOptions) ([]domain.User, uint64, error) {
	// var userModels []models.User
	// var users []entities.User
	// var count int64

	// // Executa a consulta para recuperar os produtos paginados e o total de registros correspondentes
	// result := r.Db.Offset(pagination.GetOffset()).Limit(pagination.Limit).Find(&userModels)
	// if result.Error != nil {
	// 	return nil, 0, result.Error
	// }

	// for _, userModel := range userModels {
	// 	users = append(users, *r.UserMapper.ToDomain(userModel))
	// }

	// // Executa uma consulta separada para contar o número total de registros correspondentes
	// countResult := r.Db.Model(&models.User{}).Count(&count)
	// if countResult.Error != nil {
	// 	return nil, 0, countResult.Error
	// }

	// return users, uint64(count), nil
	return nil, 0, nil
}

func (r *UserRepository) StartTransaction() error {
	return r.Base.StartTransaction()
}

func (r *UserRepository) CommitTransaction() error {
	return r.Base.CommitTransaction()
}

func (r *UserRepository) CancelTransaction() error {
	return r.Base.CancelTransaction()
}
