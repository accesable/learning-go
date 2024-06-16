// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package mysqlc

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int32
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type Item struct {
	ID               int64
	Name             string
	CategoryID       sql.NullInt32
	ShortDescription sql.NullString
	OriginalPrice    sql.NullFloat64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}