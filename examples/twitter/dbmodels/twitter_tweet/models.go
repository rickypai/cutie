// Code generated by sqlc. DO NOT EDIT.

package twitter_tweet

import (
	"database/sql"
	"encoding/json"
	"time"
)

type TwitterTweet struct {
	ID                  int64
	UserID              int64
	Text                string
	Source              string
	InReplyToStatusID   sql.NullInt64
	InReplyToUserID     sql.NullInt64
	InReplyToScreenName string
	QuotedStatusID      sql.NullInt64
	IsQuoteStatus       bool
	TwitterData         json.RawMessage
	TwitterCreatedAt    time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}