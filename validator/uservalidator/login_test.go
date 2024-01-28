package uservalidator

import (
	"game-app/param"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupValidator() Validator {
	return Validator{}
}

func TestValidate(t *testing.T) {
	defer func() { recover() }()
	v := setupValidator()
	assert.Error(t, v.ValidateLoginRequest(param.LoginRequest{}))
}
