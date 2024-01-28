package tests

import (
	"context"
	"simple-bank/src/dao"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createAccount(t *testing.T) dao.Account {
	arg := dao.CreateAccountParams{
		Owner:    "Liusha",
		Balance:  100.00,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}
