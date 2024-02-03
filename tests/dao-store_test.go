package tests

import (
	"context"
	"simple-bank/src/dao"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
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

		// TODO check accounts' balance

	}
}
