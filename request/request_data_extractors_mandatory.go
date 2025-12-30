package request

import (
	"strconv"

	"github.com/TudorHulban/epochid"
	goerrors "github.com/TudorHulban/go-errors"
)

func (data RequestData) ExtractMandatoryValue(entry string) string {
	if value, exists := data.Content[entry]; exists {
		return value
	}

	return ""
}

func (data RequestData) ExtractMandatoryEpochID(entry string) (epochid.EpochID, error) {
	if value, exists := data.Content[entry]; exists {
		result, errCr := epochid.NewEpochID(value)
		if errCr != nil {
			return 0,
				goerrors.ErrInvalidInput{
					Caller:     "ExtractMandatoryEpochID",
					Issue:      errCr,
					InputName:  "entry",
					InputValue: value,
				}
		}

		return result,
			nil
	}

	return 0,
		goerrors.ErrEntryNotFound{
			Key: entry,
		}
}

func (data RequestData) ExtractMandatoryNumber(entry string) int64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := strconv.Atoi(value)
		if errConv == nil {
			return int64(numericValue)
		}
	}

	return 0
}

func (data RequestData) ExtractMandatoryValueFromParentheses(entry string) string {
	if value, exists := data.Content[entry]; exists {
		stringValue, errConv := getValueFromParentheses(value)
		if errConv == nil {
			return stringValue
		}
	}

	return ""
}

func (data RequestData) ExtractMandatoryNumberFromParentheses(entry string) int64 {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := getNumberFromParentheses(value)
		if errConv == nil {
			return numericValue
		}
	}

	return 0
}

func (data RequestData) ExtractMandatoryValueByTags(entries ...string) string {
	for _, entry := range entries {
		if value, exists := data.Content[entry]; exists {
			return value
		}
	}

	return ""
}
