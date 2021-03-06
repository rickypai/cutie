// Code generated by sqlc. DO NOT EDIT.
// source: cpu.sql

package cpu

import (
	"context"
)

const getCPUByID = `-- name: GetCPUByID :one
SELECT id, make_id, name, cores, clock_speed_ghz, created_at, updated_at FROM cpu WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCPUByID(ctx context.Context, id int64) (CPU, error) {
	row := q.db.QueryRowContext(ctx, getCPUByID, id)
	var i CPU
	err := row.Scan(
		&i.ID,
		&i.MakeID,
		&i.Name,
		&i.Cores,
		&i.ClockSpeedGhz,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
