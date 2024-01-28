package user

import (
	"game-app/param"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	defer func() {
		recover()
	}()
	s := initService()
	_, err := s.Login(param.LoginRequest{PhoneNumber: "+989123456789", Password: "123"})
	assert.Error(t, err)
}
