package models

import (
	"github.com/axel-andrade/deu-role-auth/internal/core/domain"
)

const UserCollectionName = "users"

type UserModel struct {
	Base     BaseModel   `bson:",inline"`
	Name     string      `bson:"name"`
	Email    string      `bson:"email" validate:"required,email"`
	Password string      `bson:"password" validate:"required"`
	Role     domain.Role `bson:"role,omitempty" default:"user"`
}
