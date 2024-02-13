package tests

import (
	"context"
	"simple-bank/src/dao"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
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

func TestPassword(t *testing.T) {
	password := "secret123"
	wrongPassword := "nosecret123"
	hashedPassword, err := dao.HashPassword(password)
	require.NoError(t, err)

	err = dao.CheckPassword(hashedPassword, password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = dao.CheckPassword(hashedPassword, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
