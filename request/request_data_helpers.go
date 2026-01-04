package request

import (
	"database/sql"
	"strings"
	"time"
)

func TimeForInputDateTime(timestamp time.Time, timestampLayout string) string {
	return timestamp.Format(
		timestampLayout,
	)
}

func CleanUUID(uuid string) string {
	return strings.ReplaceAll(
		uuid,
		"-",
		"",
	)
}

func CleanOptionalUUID(uuid sql.NullString) sql.NullString {
	return sql.NullString{
		Valid: uuid.Valid,
		String: strings.ReplaceAll(
			uuid.String,
			"-",
			"",
		),
	}
}
