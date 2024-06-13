package mappers

import (
	"time"

	"github.com/axel-andrade/deu-role-auth/internal/adapters/secondary/database/mongo/models"
	"github.com/axel-andrade/deu-role-auth/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseMapper struct{}

func BuildBaseMapper() *BaseMapper {
	return &BaseMapper{}
}

func (m *BaseMapper) toDomain(model models.BaseModel) *domain.Base {
	return &domain.Base{
		ID:        model.ID.Hex(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BaseMapper) toPersistence(entity domain.Base) *models.BaseModel {
	return &models.BaseModel{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
