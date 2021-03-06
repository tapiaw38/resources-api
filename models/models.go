package models

import (
	sql "database/sql"
	"time"
)

// Null String type for sql.NullString
type NullString struct {
	sql.NullString
}

// Null Int type for sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// Null Time type for sql.NullTime
type NullTime struct {
	time.Time
	Valid bool
}

// Null Float type for sql.NullFloat64
type NullFloat64 struct {
	sql.NullFloat64
}

// Null Bool type for sql.NullBool
type NullBool struct {
	sql.NullBool
}
