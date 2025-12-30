package request

import (
	"testing"

	goerrors "github.com/TudorHulban/go-errors"
	"github.com/stretchr/testify/require"
)

func TestGetValueFromParentheses_ValidInput(t *testing.T) {
	input := "example(value)another"
	expected := "value"

	result, err := getValueFromParentheses(input)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestGetValueFromParentheses_OnlyOpeningParenthesis(t *testing.T) {
	input := "example("

	_, err := getValueFromParentheses(input)
	require.Error(t, err)

	var inputErr goerrors.ErrInvalidInput
	require.ErrorAs(t, err, &inputErr)
	require.Contains(t, err.Error(), "invalid input")
}

func TestGetValueFromParentheses_OnlyClosingParenthesis(t *testing.T) {
	input := "example)"

	_, err := getValueFromParentheses(input)
	require.Error(t, err)

	var inputErr goerrors.ErrInvalidInput
	require.ErrorAs(t, err, &inputErr)
	require.Contains(t, err.Error(), "invalid input")
}

func TestGetValueFromParentheses_NoParentheses(t *testing.T) {
	input := "example"

	_, err := getValueFromParentheses(input)
	require.Error(t, err)

	var inputErr goerrors.ErrInvalidInput
	require.ErrorAs(t, err, &inputErr)
	require.Contains(t, err.Error(), "invalid input")
}

func TestGetValueFromParentheses_MultipleParentheses(t *testing.T) {
	input := "example(first)another(second)"
	expected := "first"

	result, err := getValueFromParentheses(input)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestGetValueFromParentheses_EmptyInput(t *testing.T) {
	input := ""

	_, err := getValueFromParentheses(input)
	require.Error(t, err)

	var inputErr goerrors.ErrInvalidInput
	require.ErrorAs(t, err, &inputErr)
	require.Contains(t, err.Error(), "invalid input")
}

func TestGetValueFromParentheses_EmptyValueInParentheses(t *testing.T) {
	input := "example()"

	result, err := getValueFromParentheses(input)
	require.NoError(t, err)
	require.Empty(t, result)
}
