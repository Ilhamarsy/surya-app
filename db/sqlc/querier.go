// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	GetUser(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
}

var _ Querier = (*Queries)(nil)
