// Code generated by sqlc. DO NOT EDIT.

package twitter_user

import (
	"database/sql"
	"encoding/json"
	"time"
)

type TwitterUser struct {
	ID               int64
	ScreenName       string
	Name             string
	Description      sql.NullString
	Protected        bool
	Verified         bool
	TwitterData      json.RawMessage
	TwitterCreatedAt time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}