package request

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func (data RequestData) ExtractOptionalValue(entry string) sql.NullString {
	if value, exists := data.Content[entry]; exists {
		return sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return sql.NullString{}
}

func (data RequestData) ExtractOptionalBoolValue(entry string) sql.NullBool {
	if value, exists := data.Content[entry]; exists {
		passedValue, errConv := parseBool(value)
		if errConv == nil {
			return sql.NullBool{
				Valid: true,
				Bool:  passedValue,
			}
		}
	}

	return sql.NullBool{}
}

func (data RequestData) ExtractOptionalTimestampInLocation(entry, timestampLayout string, location *time.Location) sql.NullTime {
	if value, exists := data.Content[entry]; exists {
		parsedTime, errParse := time.ParseInLocation(timestampLayout, value, location)
		if errParse == nil {
			return sql.NullTime{
				Valid: true,
				Time:  parsedTime,
			}
		}
	}

	return sql.NullTime{}
}

func (data RequestData) ExtractOptionalNumberByTag(entry string) sql.NullInt64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := strconv.Atoi(
			strings.TrimSpace(value),
		)
		if errConv == nil {
			return sql.NullInt64{
				Valid: true,
				Int64: int64(numericValue),
			}
		}
	}

	return sql.NullInt64{}
}

func (data RequestData) ExtractOptionalNumberFromParentheses(entry string) sql.NullInt64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := getNumberFromParentheses(value)
		if errConv == nil {
			return sql.NullInt64{
				Valid: true,
				Int64: numericValue,
			}
		}
	}

	return sql.NullInt64{}
}

func (data RequestData) ExtractOptionalNumberByTags(entries ...string) sql.NullInt64 {
	for _, entry := range entries {
		if value, exists := data.Content[entry]; exists {
			numericValue, errConv := strconv.Atoi(
				strings.TrimSpace(value),
			)
			if errConv == nil {
				return sql.NullInt64{
					Valid: true,
					Int64: int64(numericValue),
				}
			}
		}
	}

	return sql.NullInt64{}
}

func (data RequestData) ExtractOptionalFloatByTag(entry string) sql.NullFloat64 {
	if value, exists := data.Content[entry]; exists {
		floatValue, errConv := strconv.ParseFloat(
			strings.TrimSpace(value),
			64,
		)
		if errConv == nil {
			return sql.NullFloat64{
				Valid:   true,
				Float64: floatValue,
			}
		}
	}

	return sql.NullFloat64{}
}
