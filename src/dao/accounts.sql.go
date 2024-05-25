// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: accounts.sql

package dao

import (
	"context"
)

const addAccountBalance = `-- name: AddAccountBalance :exec
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING id, owner, balance, currency, created_at
`

type AddAccountBalanceParams struct {
	Amount float64 `json:"amount"`
	ID     int64   `json:"id"`
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) error {
	_, err := q.db.ExecContext(ctx, addAccountBalance, arg.Amount, arg.ID)
	return err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    owner, 
    balance, 
    currency
 ) VALUES (
    $1, $2, $3
  ) 
  RETURNING id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string  `json:"owner"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, listAccounts, arg.Owner, arg.Limit, arg.Offset)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING id, owner, balance, currency, created_at
`

type UpdateAccountParams struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.ExecContext(ctx, updateAccount, arg.ID, arg.Balance)
	return err
}
