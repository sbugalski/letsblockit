// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type BannedUser struct {
	ID         int32
	UserID     uuid.UUID
	CreatedAt  time.Time
	Reason     string
	LiftedAt   sql.NullTime
	LiftReason sql.NullString
}

type FilterInstance struct {
	ID           int32
	UserID       uuid.UUID
	FilterListID int32
	FilterName   string
	Params       pgtype.JSONB
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}

type FilterList struct {
	ID           int32
	UserID       uuid.UUID
	Token        uuid.UUID
	CreatedAt    time.Time
	Downloaded   bool
	DownloadedAt sql.NullTime
}
