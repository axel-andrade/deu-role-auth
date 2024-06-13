package value_object

import (
	"fmt"

	err_msg "github.com/axel-andrade/deu-role-auth/internal/core/domain/constants/errors"
)

type Password struct {
	Value string
}

func (p *Password) Validate() error {
	if length := len(p.Value); length >= 6 {
		return nil
	}

	return fmt.Errorf(err_msg.INVALID_PASSWORD)
}
