package user_test

import (
	"testing"

	"github.com/gui-laranjeira/testyfy/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type test struct {
		name          string
		email         string
		password      string
		phone         string
		expectedError bool
	}
	tests := []test{
		{
			name:          "valid user",
			email:         "valid@mail.com",
			password:      "ValidPassword1!",
			phone:         "123456789",
			expectedError: false,
		},
		{
			name:          "",
			email:         "valid@mail.com",
			password:      "ValidPassword1!",
			phone:         "123456789",
			expectedError: true,
		},
		{
			name:          "valid user",
			email:         "",
			password:      "ValidPassword1!",
			phone:         "123456789",
			expectedError: true,
		},
		{
			name:          "valid user",
			email:         "valid@mail.com",
			password:      "",
			phone:         "123456789",
			expectedError: true,
		},
		{
			name:          "valid user",
			email:         "valid@mail.com",
			password:      "ValidPassword1!",
			phone:         "",
			expectedError: true,
		},
		{
			name:          "valid user",
			email:         "valid@mail.com",
			password:      "invalid",
			phone:         "123456789",
			expectedError: true,
		},
	}

	for _, tc := range tests {
		_, err := user.NewUser(tc.email, tc.password, tc.name, tc.phone)
		assert.Equal(t, tc.expectedError, err != nil)
	}
}

func TestComparePassword(t *testing.T) {
	u, _ := user.NewUser("test@test.com", "ValidPassword1!", "hello", "123456789")
	assert.True(t, u.ComparePassword("ValidPassword1!"))
}
