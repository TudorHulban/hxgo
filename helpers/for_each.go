package helpers

import "testing"

func ForEachValue[T, R any](values []T, process func(T) R) []R {
	result := make([]R, len(values), len(values)) //nolint:staticcheck

	for ix, value := range values {
		result[ix] = process(value)
	}

	return result
}

func ForEachTest[T any](t *testing.T, values []T, process func(T, *testing.T)) {
	t.Helper()

	for _, value := range values {
		process(value, t)
	}
}

type ParamsForEachValueWAddition[T, R any] struct {
	Addition func() R
	Process  func(T) R

	Values []T
}

func ForEachValueWAddition[T, R any](params *ParamsForEachValueWAddition[T, R]) []R {
	if len(params.Values) == 0 {
		return nil
	}

	result := make([]R, 0, len(params.Values)+1)

	result = append(result, params.Addition())

	for _, value := range params.Values {
		result = append(result, params.Process(value))
	}

	return result
}
