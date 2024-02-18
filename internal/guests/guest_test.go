package guests_test

import (
	"testing"

	"github.com/gui-laranjeira/testyfy/internal/guests"
	"github.com/stretchr/testify/assert"
)

func TestNewGuest(t *testing.T) {
	type test struct {
		name          string
		userId        string
		email         string
		expectedError bool
	}

	tests := []test{
		{
			name:          "valid guest",
			userId:        "4f6e6cf8-0361-4f3a-8c4e-1f0e4e0e3c3b",
			email:         "test@test.com",
			expectedError: false,
		},
		{
			name:          "invalid user id",
			userId:        "invalid",
			email:         "test@test.com",
			expectedError: true,
		},
		{
			name:          "invalid email",
			userId:        "4f6e6cf8-0361-4f3a-8c4e-1f0e4e0e3c3b",
			email:         "invalid",
			expectedError: true,
		},
		{
			name:          "empty email",
			userId:        "4f6e6cf8-0361-4f3a-8c4e-1f0e4e0e3c3b",
			email:         "",
			expectedError: true,
		},
		{
			name:          "",
			userId:        "4f6e6cf8-0361-4f3a-8c4e-1f0e4e0e3c3b",
			email:         "test@test.com",
			expectedError: true,
		},
	}
	for _, tc := range tests {
		_, err := guests.NewGuest(tc.userId, tc.name, tc.email)
		assert.Equal(t, tc.expectedError, err != nil)
	}

}
