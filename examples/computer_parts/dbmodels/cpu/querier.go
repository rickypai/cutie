// Code generated by sqlc. DO NOT EDIT.

package cpu

import (
	"context"
)

type Querier interface {
	GetCPUByID(ctx context.Context, id int64) (CPU, error)
}

var _ Querier = (*Queries)(nil)
