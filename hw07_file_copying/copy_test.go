package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	const from = "testdata/input.txt"

	t.Run("limit 0 offset 0", func(t *testing.T) {
		_, err := os.CreateTemp("", "out_offset0_limit0.txt")
		require.NoError(t, err)

		to := "out_offset0_limit0.txt"
		targetTest := "testdata/out_offset0_limit0.txt"
		var limit, offset int64

		err = Copy(from, to, limit, offset)
		require.NoError(t, err)

		testBytes, err := os.ReadFile(targetTest)
		require.NoError(t, err)

		targetBytes, err := os.ReadFile(to)
		require.NoError(t, err)

		require.Equal(t, testBytes, targetBytes)
	})
}

func TestCopyLimit(t *testing.T) {
	const from = "testdata/input.txt"
	const directory = "testdata/"

	tests := []struct {
		name, to, targetTest string
		limit, offset        int64
	}{
		{name: "limit 10", to: "limit10.txt", targetTest: "out_offset0_limit10.txt", limit: 10, offset: 0},
		{name: "limit 1000", to: "limit1000.txt", targetTest: "out_offset0_limit1000.txt", limit: 1000, offset: 0},
		{name: "limit 10000", to: "limit10000.txt", targetTest: "out_offset0_limit10000.txt", limit: 10000, offset: 0},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := os.CreateTemp("", tc.to)
			require.NoError(t, err)

			err = Copy(from, tc.to, tc.limit, tc.offset)
			require.NoError(t, err)

			testFile := fmt.Sprintf("%s%s", directory, tc.targetTest)
			testBytes, err := os.ReadFile(testFile)
			require.NoError(t, err)

			targetBytes, err := os.ReadFile(tc.to)
			require.NoError(t, err)

			require.Equal(t, testBytes, targetBytes)
		})
	}
}

func TestCopyLimitAndOffset(t *testing.T) {
	const from = "testdata/input.txt"
	const directory = "testdata/"

	tests := []struct {
		name, to, targetTest string
		limit, offset        int64
	}{
		{name: "offset 100", to: "offset100.txt", targetTest: "out_offset100_limit1000.txt", limit: 1000, offset: 100},
		{name: "offset 6000", to: "offset6000.txt", targetTest: "out_offset6000_limit1000.txt", limit: 1000, offset: 6000},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := os.CreateTemp("", tc.to)
			require.NoError(t, err)

			err = Copy(from, tc.to, tc.limit, tc.offset)
			require.NoError(t, err)

			testFile := fmt.Sprintf("%s%s", directory, tc.targetTest)
			testBytes, err := os.ReadFile(testFile)
			require.NoError(t, err)

			targetBytes, err := os.ReadFile(tc.to)
			require.NoError(t, err)

			require.Equal(t, testBytes, targetBytes)
		})
	}
}

func TestCopyError(t *testing.T) {
	t.Run("offset exceeds file size", func(t *testing.T) {
		_, err := os.CreateTemp("", "out_offset_exceeds.txt")
		require.NoError(t, err)

		from := "testdata/input.txt"
		to := "out_offset_exceeds.txt"
		var limit int64
		var offset int64 = 7000

		err = Copy(from, to, limit, offset)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize, "actual err - %v")
	})
}
