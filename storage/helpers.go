package storage

import "database/sql"

// helper function to control the null fields

// StringToNull is a function to convert a string to a sql.NullString
func StringToNull(str string) sql.NullString {
	null := sql.NullString{String: str}
	if null.String == "" {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}

// IntToNull is a function to convert a int to a sql.NullInt64
func IntToNull(i int64) sql.NullInt64 {
	null := sql.NullInt64{Int64: i}
	if null.Int64 == 0 {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}

// FloatToNull is a function to convert a float to a sql.NullFloat64
func FloatToNull(f float64) sql.NullFloat64 {
	null := sql.NullFloat64{Float64: f}
	if null.Float64 == 0 {
		null.Valid = false
	} else {
		null.Valid = true
	}
	return null
}
