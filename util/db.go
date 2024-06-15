package util

import (
	"database/sql"
	"time"
)

func NullToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func NullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func NullToBool(s sql.NullBool) bool {
	if s.Valid {
		return s.Bool
	}
	return false
}

func NullBool(s bool) sql.NullBool {
	return sql.NullBool{
		Bool:  s,
		Valid: true,
	}
}

func NullToTime(s sql.NullTime) time.Time {
	if s.Valid {
		return s.Time
	}
	return time.Time{}
}

func NullTime(s time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  s,
		Valid: !s.IsZero(),
	}
}

func NullToInt64(s sql.NullInt64) int64 {
	if s.Valid {
		return s.Int64
	}
	return 0
}

func NullInt64(s int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: s,
		Valid: true,
	}
}
