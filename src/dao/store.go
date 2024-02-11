package dao

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	TransferTx(context.Context, TransferTxParams) (TransferTxResult, error)
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

type QueryHandler func(*Queries) error

type TransferTxParams struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within database transaction
func (store *SQLStore) execTx(ctx context.Context, fn QueryHandler) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err : %v, rollback err : %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func updateAccounts(
	ctx context.Context,
	q *Queries,
	account1_id int64,
	amount1 float64,
	account2_id int64,
	amount2 float64,
) error {
	err := q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     account1_id,
		Amount: amount1,
	})
	if err != nil {
		return err
	}

	err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     account2_id,
		Amount: amount2,
	})
	if err != nil {
		return err
	}

	return nil
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx,
		func(q *Queries) error {
			var err error

			result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
				FromAccountID: arg.FromAccountID,
				ToAccountID:   arg.ToAccountID,
				Amount:        arg.Amount,
			})
			if err != nil {
				return err
			}

			result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
				AccountID: arg.FromAccountID,
				Amount:    -arg.Amount,
			})
			if err != nil {
				return err
			}

			result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
				AccountID: arg.ToAccountID,
				Amount:    arg.Amount,
			})
			if err != nil {
				return err
			}

			if arg.FromAccountID < arg.ToAccountID {
				err = updateAccounts(
					ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount,
				)
			} else {
				err = updateAccounts(
					ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount,
				)
			}

			return err
		},
	)

	return result, err
}
