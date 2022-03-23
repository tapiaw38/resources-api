package models

import "database/sql"

// Null String type for sql.NullString
type NullString struct {
	sql.NullString
}

// Null Int type for sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

type NullInt32 struct {
	sql.NullInt32
}

// Null Date type for sql.NullTime
type NullDate struct {
	sql.NullString
}

// Null Time type for sql.NullTime
type NullTime struct {
	sql.NullTime
}

// Null Float type for sql.NullFloat64
type NullFloat struct {
	sql.NullFloat64
}

// Null Bool type for sql.NullBool
type NullBool struct {
	sql.NullBool
}
