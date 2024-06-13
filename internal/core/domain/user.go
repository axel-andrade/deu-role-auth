package domain

import (
	"fmt"

	vo "github.com/axel-andrade/deu-role-auth/internal/core/domain/value_objects"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Base
	Name     vo.Name     `json:"name"`
	Email    vo.Email    `json:"email"`
	Password vo.Password `json:"-"`
	Role     Role        `json:"role"`
}

func BuildUser(name, email, password string, role Role) (*User, error) {
	u := &User{
		Name:     vo.Name{Value: name},
		Email:    vo.Email{Value: email},
		Password: vo.Password{Value: password},
		Role:     Role(role),
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) validate() error {
	if err := u.Name.Validate(); err != nil {
		return err
	}

	if err := u.Email.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	if u.Role != RoleAdmin && u.Role != RoleUser {
		return fmt.Errorf("invalid role")
	}

	return nil
}
