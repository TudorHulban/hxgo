package request

import (
	"testing"

	"github.com/TudorHulban/epochid"
	"github.com/stretchr/testify/require"
)

func TestValueDifferentOrMissing(t *testing.T) {
	data := RequestData{
		Content: map[string]string{
			"test_key": "(123)",
		},
	}

	require.False(t,
		data.ValueDifferentOrMissing("test_key", 123),
		"T1. Values match, should NOT be different",
	)

	require.True(t,
		data.ValueDifferentOrMissing("test_key", 124),
		"T2. Value differs, should be different",
	)

	delete(data.Content, "test_key")
	require.True(t,
		data.ValueDifferentOrMissing("test_key", 123),
		"T3. Entry missing",
	)

	data.Content["test_key"] = "abc"
	require.True(t,
		data.ValueDifferentOrMissing("test_key", 123),
		"T4. Invalid format, should be different",
	)
}

func TestEpochNotInSlice(t *testing.T) {
	data1 := RequestData{Content: map[string]string{}}
	slice1 := []epochid.EpochID{1737541841100022188}
	require.True(t,
		data1.EntryValueNotInSlice("key1", slice1),
		"T1: Empty map",
	)

	data2 := RequestData{Content: map[string]string{"key1": "(1737541841100022188)"}}
	slice2 := []epochid.EpochID{1737541841100022188}
	require.True(t,
		data2.EntryValueNotInSlice("key2", slice2),
		"T2: Key not in map",
	)

	data3 := RequestData{Content: map[string]string{"key1": "(1737541841100022188)"}}
	slice3 := []epochid.EpochID{1737541841100022188}
	require.False(t,
		data3.EntryValueNotInSlice("key1", slice3),
		"T3: Value in slice",
	)

	data4 := RequestData{Content: map[string]string{"key1": "(1737541841100022188)"}}
	slice4 := []epochid.EpochID{1737541841100022100}
	require.True(t,
		data4.EntryValueNotInSlice("key1", slice4),
		"T4: Value not in slice",
	)

	data5 := RequestData{Content: map[string]string{"key1": "abc"}}
	slice5 := []epochid.EpochID{1737541841100022188}
	require.True(t,
		data5.EntryValueNotInSlice("key1", slice5),
		"T5: Invalid format",
	)
}
