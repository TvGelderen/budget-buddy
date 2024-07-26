// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transactions.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated
`

type CreateTransactionParams struct {
	UserID       uuid.UUID
	Amount       string
	Description  string
	Incoming     bool
	Type         string
	Recurring    bool
	StartDate    time.Time
	EndDate      time.Time
	Interval     sql.NullString
	DaysInterval sql.NullInt32
	Created      time.Time
	Updated      time.Time
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.UserID,
		arg.Amount,
		arg.Description,
		arg.Incoming,
		arg.Type,
		arg.Recurring,
		arg.StartDate,
		arg.EndDate,
		arg.Interval,
		arg.DaysInterval,
		arg.Created,
		arg.Updated,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Description,
		&i.Incoming,
		&i.Type,
		&i.Recurring,
		&i.StartDate,
		&i.EndDate,
		&i.Interval,
		&i.DaysInterval,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions 
WHERE id = $1 AND user_id = $2
`

type DeleteTransactionParams struct {
	ID     int64
	UserID uuid.UUID
}

func (q *Queries) DeleteTransaction(ctx context.Context, arg DeleteTransactionParams) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, arg.ID, arg.UserID)
	return err
}

const getUserExpenseTransactionsBetweenDates = `-- name: GetUserExpenseTransactionsBetweenDates :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1 AND NOT incoming AND start_date <= $2 AND end_date >= $3
ORDER BY start_date
LIMIT $4
OFFSET $5
`

type GetUserExpenseTransactionsBetweenDatesParams struct {
	UserID    uuid.UUID
	StartDate time.Time
	EndDate   time.Time
	Limit     int32
	Offset    int32
}

func (q *Queries) GetUserExpenseTransactionsBetweenDates(ctx context.Context, arg GetUserExpenseTransactionsBetweenDatesParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserExpenseTransactionsBetweenDates,
		arg.UserID,
		arg.StartDate,
		arg.EndDate,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserIncomeTransactionsBetweenDates = `-- name: GetUserIncomeTransactionsBetweenDates :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1 AND incoming AND start_date <= $2 AND end_date >= $3
ORDER BY start_date
LIMIT $4
OFFSET $5
`

type GetUserIncomeTransactionsBetweenDatesParams struct {
	UserID    uuid.UUID
	StartDate time.Time
	EndDate   time.Time
	Limit     int32
	Offset    int32
}

func (q *Queries) GetUserIncomeTransactionsBetweenDates(ctx context.Context, arg GetUserIncomeTransactionsBetweenDatesParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserIncomeTransactionsBetweenDates,
		arg.UserID,
		arg.StartDate,
		arg.EndDate,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserIncomingTransactions = `-- name: GetUserIncomingTransactions :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1 AND incoming = 1
`

func (q *Queries) GetUserIncomingTransactions(ctx context.Context, userID uuid.UUID) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserIncomingTransactions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserOutgoingTransactions = `-- name: GetUserOutgoingTransactions :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1 AND incoming = 0
`

func (q *Queries) GetUserOutgoingTransactions(ctx context.Context, userID uuid.UUID) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserOutgoingTransactions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTransactions = `-- name: GetUserTransactions :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1
ORDER BY start_date
LIMIT $2
OFFSET $3
`

type GetUserTransactionsParams struct {
	UserID uuid.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) GetUserTransactions(ctx context.Context, arg GetUserTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserTransactions, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTransactionsBetweenDates = `-- name: GetUserTransactionsBetweenDates :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1 AND start_date <= $2 AND end_date >= $3
ORDER BY start_date
LIMIT $4
OFFSET $5
`

type GetUserTransactionsBetweenDatesParams struct {
	UserID    uuid.UUID
	StartDate time.Time
	EndDate   time.Time
	Limit     int32
	Offset    int32
}

func (q *Queries) GetUserTransactionsBetweenDates(ctx context.Context, arg GetUserTransactionsBetweenDatesParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserTransactionsBetweenDates,
		arg.UserID,
		arg.StartDate,
		arg.EndDate,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTransactionsByType = `-- name: GetUserTransactionsByType :many
SELECT id, user_id, amount, description, incoming, type, recurring, start_date, end_date, interval, days_interval, created, updated FROM transactions 
WHERE user_id = $1 AND type = $2
`

type GetUserTransactionsByTypeParams struct {
	UserID uuid.UUID
	Type   string
}

func (q *Queries) GetUserTransactionsByType(ctx context.Context, arg GetUserTransactionsByTypeParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserTransactionsByType, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Description,
			&i.Incoming,
			&i.Type,
			&i.Recurring,
			&i.StartDate,
			&i.EndDate,
			&i.Interval,
			&i.DaysInterval,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransaction = `-- name: UpdateTransaction :exec
UPDATE transactions
SET amount = $3, description = $4, incoming = $5, type = $6, recurring = $7, start_date = $8, end_date = $9, interval = $10, days_interval = $11, updated = $12
WHERE id = $1 AND user_id = $2
`

type UpdateTransactionParams struct {
	ID           int64
	UserID       uuid.UUID
	Amount       string
	Description  string
	Incoming     bool
	Type         string
	Recurring    bool
	StartDate    time.Time
	EndDate      time.Time
	Interval     sql.NullString
	DaysInterval sql.NullInt32
	Updated      time.Time
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, updateTransaction,
		arg.ID,
		arg.UserID,
		arg.Amount,
		arg.Description,
		arg.Incoming,
		arg.Type,
		arg.Recurring,
		arg.StartDate,
		arg.EndDate,
		arg.Interval,
		arg.DaysInterval,
		arg.Updated,
	)
	return err
}
