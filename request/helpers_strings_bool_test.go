package request

import (
	"testing"

	goerrors "github.com/TudorHulban/go-errors"
	"github.com/stretchr/testify/require"
)

func TestParseBool(t *testing.T) {
	tests := []struct {
		description string
		input       string

		expectedValue      bool
		expectUnknownValue bool
	}{
		// error cases first
		{
			description:        "1_empty_string",
			input:              "",
			expectUnknownValue: true,
		},
		{
			description:        "2_random_string",
			input:              "abc",
			expectUnknownValue: true,
		},
		{
			description:        "3_number_not_allowed",
			input:              "2",
			expectUnknownValue: true,
		},
		{
			description:        "4_whitespace",
			input:              " true ",
			expectUnknownValue: true,
		},

		// valid false values
		{
			description:   "5_false_lowercase",
			input:         "false",
			expectedValue: false,
		},
		{
			description:   "6_false_uppercase",
			input:         "FALSE",
			expectedValue: false,
		},
		{
			description:   "7_false_mixedcase",
			input:         "False",
			expectedValue: false,
		},
		{
			description:   "8_no",
			input:         "no",
			expectedValue: false,
		},
		{
			description:   "9_f_lowercase",
			input:         "f",
			expectedValue: false,
		},
		{
			description:   "10_F_uppercase",
			input:         "F",
			expectedValue: false,
		},
		{
			description:   "11_zero",
			input:         "0",
			expectedValue: false,
		},

		// valid true values
		{
			description:   "12_true_lowercase",
			input:         "true",
			expectedValue: true,
		},
		{
			description:   "13_true_uppercase",
			input:         "TRUE",
			expectedValue: true,
		},
		{
			description:   "14_true_mixedcase",
			input:         "True",
			expectedValue: true,
		},
		{
			description:   "15_yes",
			input:         "yes",
			expectedValue: true,
		},
		{
			description:   "16_t_lowercase",
			input:         "t",
			expectedValue: true,
		},
		{
			description:   "17_T_uppercase",
			input:         "T",
			expectedValue: true,
		},
		{
			description:   "18_one",
			input:         "1",
			expectedValue: true,
		},
	}

	for _, tc := range tests {
		t.Run(
			tc.description,
			func(t *testing.T) {
				result, err := parseBool(tc.input)

				if tc.expectUnknownValue {
					require.Error(t, err)
					require.ErrorIs(t, err, goerrors.ErrUnknownValue)
					return
				}

				require.NoError(t, err)
				require.Equal(t, tc.expectedValue, result)
			},
		)
	}
}

// BenchmarkParseBool/1_empty_string-16         	657321590	         1.919 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/2_random_string-16        	343580024	         3.235 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/3_number_not_allowed-16   	382334838	         3.098 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/4_whitespace-16           	278651815	         3.823 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/5_false_lowercase-16      	161504205	         6.632 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/6_false_uppercase-16      	364441042	         3.450 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/7_no-16                   	365788549	         3.629 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/8_zero-16                 	298417923	         4.021 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/9_true_lowercase-16       	271671697	         4.662 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/10_true_uppercase-16      	420758496	         2.799 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/11_yes-16                 	234551835	         4.914 ns/op	       0 B/op	       0 allocs/op
// BenchmarkParseBool/12_one-16                 	397391175	         4.441 ns/op	       0 B/op	       0 allocs/op

func BenchmarkParseBool(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		// error cases first
		{
			name:  "1_empty_string",
			input: "",
		},
		{
			name:  "2_random_string",
			input: "abc",
		},
		{
			name:  "3_number_not_allowed",
			input: "2",
		},
		{
			name:  "4_whitespace",
			input: " true ",
		},

		// valid false values
		{
			name:  "5_false_lowercase",
			input: "false",
		},
		{
			name:  "6_false_uppercase",
			input: "FALSE",
		},
		{
			name:  "7_no",
			input: "no",
		},
		{
			name:  "8_zero",
			input: "0",
		},

		// valid true values
		{
			name:  "9_true_lowercase",
			input: "true",
		},
		{
			name:  "10_true_uppercase",
			input: "TRUE",
		},
		{
			name:  "11_yes",
			input: "yes",
		},
		{
			name:  "12_one",
			input: "1",
		},
	}

	for _, tc := range tests {
		b.Run(
			tc.name,
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = parseBool(tc.input)
				}
			},
		)
	}
}
