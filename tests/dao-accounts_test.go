package tests

import (
	"context"
	"simple-bank/src/dao"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createAccount(arg dao.CreateAccountParams) (dao.Account, error) {
	return testQueries.CreateAccount(context.Background(), arg)
}

func TestCreateAccount(t *testing.T) {
	user, err := createUser(dao.CreateUserParams{
		Username:       "test1",
		HashedPassword: "secret",
		Email:          "test1@test.org",
		FullName:       "test one",
	})

	arg := dao.CreateAccountParams{
		Owner:    user.Username,
		Balance:  1500.00,
		Currency: "USD",
	}

	account, err := createAccount(arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	user, err := createUser(dao.CreateUserParams{
		Username:       "test2",
		HashedPassword: "secret",
		Email:          "test2@test.org",
		FullName:       "test two",
	})

	arg := dao.CreateAccountParams{
		Owner:    user.Username,
		Balance:  1500.00,
		Currency: "GBP",
	}

	account1, err := createAccount(arg)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}
