package request

import (
	"slices"

	"github.com/TudorHulban/epochid"
)

func (data RequestData) ValueDifferentOrMissing(entry string, compareWith int64) bool {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := getNumberFromParentheses(value)
		if errConv != nil {
			return true
		}

		return numericValue != compareWith
	}

	return true
}

func (data RequestData) EntryValueNotInSlice(entry string, slice []epochid.EpochID) bool {
	if value, exists := data.Content[entry]; exists {
		numericValue, errConv := getNumberFromParentheses(value)
		if errConv != nil {
			return true
		}

		if slices.Index(slice, epochid.EpochID(numericValue)) != -1 {
			return false
		}
	}

	return true
}
