package mappers

import (
	"github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo/models"
	"github.com/axel-andrade/deu-role-auth/internal/core/domain"
	value_object "github.com/axel-andrade/deu-role-auth/internal/core/domain/value_objects"
)

type UserMapper struct {
	BaseMapper
}

func BuildUserMapper(baseMapper *BaseMapper) *UserMapper {
	return &UserMapper{BaseMapper: *baseMapper}
}

func (m *UserMapper) ToDomain(model models.UserModel) *domain.User {
	return &domain.User{
		Base:     *m.BaseMapper.toDomain(model.Base),
		Email:    value_object.Email{Value: model.Email},
		Name:     value_object.Name{Value: model.Name},
		Password: value_object.Password{Value: model.Password},
		Role:     domain.Role(model.Role),
	}
}

func (m *UserMapper) ToPersistence(entity domain.User) models.UserModel {
	return models.UserModel{
		Base:     *m.BaseMapper.toPersistence(entity.Base),
		Email:    entity.Email.Value,
		Name:     entity.Name.Value,
		Password: entity.Password.Value,
		Role:     entity.Role,
	}
}
