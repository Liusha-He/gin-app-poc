package tests

import (
	"context"
	"fmt"
	"simple-bank/src/dao"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	// existed := make(map[int]bool)
	store := dao.NewStore(testDB)

	account1, _ := testQueries.CreateAccount(context.Background(),
		dao.CreateAccountParams{
			Owner:    "Liusha",
			Balance:  20000.00,
			Currency: "GBP",
		},
	)
	account2, _ := testQueries.CreateAccount(context.Background(),
		dao.CreateAccountParams{
			Owner:    "degere",
			Balance:  10000.00,
			Currency: "GBP",
		},
	)
	fmt.Printf(
		"Before update, liusha's balance is £%.2f : Degere's balance is £%.2f\n",
		account1.Balance,
		account2.Balance,
	)

	n := 5
	amount := 300.00

	errs := make(chan error)
	results := make(chan dao.TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), dao.TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
		require.Equal(t, transfer.Amount, amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.AccountID, account1.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		require.Equal(t, -amount, fromEntry.Amount)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.AccountID, account2.ID)
		require.NotZero(t, toEntry.CreatedAt)
		require.Equal(t, amount, toEntry.Amount)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
	}

	// check the final updated accounts
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	diff := amount * float64(n)
	require.Equal(t, account1.Balance-diff, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+diff, updatedAccount2.Balance)

	fmt.Printf(
		"After transaction liusha's balance is £%.2f and Degere's balance is £%.2f",
		updatedAccount1.Balance,
		updatedAccount2.Balance,
	)
}

func TestTransferTxDeadLock(t *testing.T) {
	// existed := make(map[int]bool)
	store := dao.NewStore(testDB)

	account1, _ := testQueries.CreateAccount(context.Background(),
		dao.CreateAccountParams{
			Owner:    "Liusha",
			Balance:  20000.00,
			Currency: "GBP",
		},
	)
	account2, _ := testQueries.CreateAccount(context.Background(),
		dao.CreateAccountParams{
			Owner:    "degere",
			Balance:  10000.00,
			Currency: "GBP",
		},
	)

	n := 10
	amount := 300.00

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountId := account1.ID
		toAccountId := account2.ID

		if i%2 == 0 {
			fromAccountId = account2.ID
			toAccountId = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), dao.TransferTxParams{
				FromAccountID: fromAccountId,
				ToAccountID:   toAccountId,
				Amount:        amount,
			})

			errs <- err

		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}

	// check the final updated accounts
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
