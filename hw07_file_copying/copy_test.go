package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	const from = "testdata/input.txt"

	t.Run("limit 0 offset 0", func(t *testing.T) {
		to := "tmp/out_offset0_limit0.txt"
		targetTest := "testdata/out_offset0_limit0.txt"
		var limit, offset int64

		err := Copy(from, to, offset, limit)
		if err != nil {
			t.Fatalf("error on copy file")
		}

		testBytes, err := os.ReadFile(targetTest)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		targetBytes, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		if string(testBytes) != string(targetBytes) {
			t.Fatalf("bad data in target file after copy")
		}
	})
}

func TestCopyLimit(t *testing.T) {
	const from = "testdata/input.txt"

	t.Run("limit 10 offset 0", func(t *testing.T) {
		to := "tmp/out_offset0_limit10.txt"
		targetTest := "testdata/out_offset0_limit10.txt"
		var limit int64 = 10
		var offset int64

		err := Copy(from, to, offset, limit)
		if err != nil {
			t.Fatalf("error on copy file")
		}

		testBytes, err := os.ReadFile(targetTest)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		targetBytes, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		if string(testBytes) != string(targetBytes) {
			t.Fatalf("bad data in target file after copy")
		}
	})

	t.Run("limit 1000 offset 0", func(t *testing.T) {
		to := "tmp/out_offset0_limit1000.txt"
		targetTest := "testdata/out_offset0_limit1000.txt"
		var limit int64 = 1000
		var offset int64

		err := Copy(from, to, offset, limit)
		if err != nil {
			t.Fatalf("error on copy file")
		}

		testBytes, err := os.ReadFile(targetTest)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		targetBytes, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		if string(testBytes) != string(targetBytes) {
			t.Fatalf("bad data in target file after copy")
		}
	})

	t.Run("limit 10000 offset 0", func(t *testing.T) {
		to := "tmp/out_offset0_limit10000.txt"
		targetTest := "testdata/out_offset0_limit10000.txt"
		var limit int64 = 10000
		var offset int64

		err := Copy(from, to, offset, limit)
		if err != nil {
			t.Fatalf("error on copy file %v", err)
		}

		testBytes, err := os.ReadFile(targetTest)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		targetBytes, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		if string(testBytes) != string(targetBytes) {
			t.Fatalf("bad data in target file after copy")
		}
	})
}

func TestCopyLimitAndOffset(t *testing.T) {
	const from = "testdata/input.txt"

	t.Run("limit 1000 offset 100", func(t *testing.T) {
		to := "tmp/out_offset100_limit1000.txt"
		targetTest := "testdata/out_offset100_limit1000.txt"
		var limit int64 = 1000
		var offset int64 = 100

		err := Copy(from, to, offset, limit)
		if err != nil {
			t.Fatalf("error on copy file")
		}

		testBytes, err := os.ReadFile(targetTest)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		targetBytes, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		if string(testBytes) != string(targetBytes) {
			t.Fatalf("bad data in target file after copy")
		}
	})

	t.Run("limit 1000 offset 6000", func(t *testing.T) {
		to := "tmp/out_offset6000_limit1000.txt"
		targetTest := "testdata/out_offset6000_limit1000.txt"
		var limit int64 = 1000
		var offset int64 = 6000

		err := Copy(from, to, offset, limit)
		if err != nil {
			t.Fatalf("error on copy file")
		}

		testBytes, err := os.ReadFile(targetTest)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		targetBytes, err := os.ReadFile(to)
		if err != nil {
			t.Fatalf("error on read target file")
		}

		if string(testBytes) != string(targetBytes) {
			t.Fatalf("bad data in target file after copy")
		}
	})
}

func TestCopyError(t *testing.T) {
	t.Run("offset exceeds file size", func(t *testing.T) {
		from := "testdata/input.txt"
		to := "tmp/out_offset_exceeds.txt"
		var limit int64
		var offset int64 = 7000

		err := Copy(from, to, offset, limit)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize, "actual err - %v")
	})

	t.Run("copy directory", func(t *testing.T) {
		from := "/testdata"
		to := "tmp/out_not_open_file.txt"
		var limit int64
		var offset int64 = 7000

		err := Copy(from, to, offset, limit)
		require.ErrorIs(t, err, ErrUnsupportedFile, "actual err - %v")
	})
}
