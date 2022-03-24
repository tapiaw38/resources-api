package models

import (
	sql "database/sql"
)

// Null String type for sql.NullString
type NullString struct {
	sql.NullString
}

// Null Int type for sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// Null Float type for sql.NullFloat64
type NullFloat struct {
	sql.NullFloat64
}

// Null Bool type for sql.NullBool
type NullBool struct {
	sql.NullBool
}
