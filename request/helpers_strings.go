package request

import (
	"errors"
	"strconv"
	"strings"

	goerrors "github.com/TudorHulban/go-errors"
)

func getValueFromParentheses(text string) (string, error) {
	_, after, existsLeftP := strings.Cut(text, "(")
	if existsLeftP {
		before, _, existsRightP := strings.Cut(after, ")")
		if existsRightP {
			return before,
				nil
		}
	}

	return "",
		goerrors.ErrInvalidInput{
			Caller:     "getValueFromParentheses",
			InputName:  "value",
			InputValue: text,
			Issue:      errors.New("invalid input"),
		}
}

func getNumberFromParentheses(text string) (int64, error) {
	_, after, existsLeftP := strings.Cut(text, "(")
	if existsLeftP {
		before, _, existsRightP := strings.Cut(after, ")")
		if existsRightP {
			return strconv.ParseInt(before, 10, 64)
		}
	}

	return 0,
		goerrors.ErrInvalidInput{
			Caller:     "getNumberFromParentheses",
			InputName:  "number",
			InputValue: text,
			Issue:      errors.New("invalid input"),
		}
}

// Use TRUE or FALSE for fastest match.
// See benchmark.
func parseBool(value string) (bool, error) {
	switch value {
	case "1", "t", "T", "true", "TRUE", "True", "yes":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "no":
		return false, nil
	}

	return false,
		goerrors.ErrUnknownValue
}
