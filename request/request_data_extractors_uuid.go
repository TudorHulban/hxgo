package request

import (
	"database/sql"
	"strconv"
)

func (data RequestData) ExtractMandatoryUUID(entry string) int64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := strconv.Atoi(
			CleanUUID(value),
		)
		if errConv == nil {
			return int64(numericValue)
		}
	}

	return 0
}

func (data RequestData) ExtractOptionalUUID(entry string) sql.NullInt64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := strconv.Atoi(
			CleanUUID(value),
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

func (data RequestData) ExtractMandatoryUUIDFromParentheses(entry string) int64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := getNumberFromParentheses(
			CleanUUID(value),
		)
		if errConv == nil {
			return numericValue
		}
	}

	return 0
}

func (data RequestData) ExtractMandatoryUUIDFromParenthesesByTags(entries ...string) int64 {
	for _, entry := range entries {
		if value, exists := data.Content[entry]; exists {
			numericValue, errConv := getNumberFromParentheses(
				CleanUUID(value),
			)
			if errConv == nil {
				return numericValue
			}
		}
	}

	return 0
}
