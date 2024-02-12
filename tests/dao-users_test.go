package tests

import (
	"context"
	"simple-bank/src/dao"
	"testing"

	"github.com/stretchr/testify/require"
)

func createUser(arg dao.CreateUserParams) (dao.User, error) {
	return testQueries.CreateUser(context.Background(), arg)
}

func TestCreateUser(t *testing.T) {
	arg := dao.CreateUserParams{
		Username:       "user",
		HashedPassword: "secret",
		FullName:       "test user",
		Email:          "user@test.com",
	}

	user, err := createUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)
}
