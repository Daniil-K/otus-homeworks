package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadTestDir(t *testing.T) {
	testdata := Environment{
		"BAR":   EnvValue{"bar", false},
		"EMPTY": EnvValue{"", false},
		"FOO":   EnvValue{"   foo\nwith new line", false},
		"HELLO": EnvValue{"\"hello\"", false},
		"UNSET": EnvValue{"", true},
	}
	result, err := ReadDir("testdata/env")
	fmt.Println(result)
	fmt.Println(testdata)
	require.NoError(t, err)
	require.Equal(t, result, testdata)
}

func TestEmptyDir(t *testing.T) {
	tmp := t.TempDir()

	defer func() {
		err := os.Remove(tmp)
		if err != nil {
			log.Fatal(err)
		}
	}()

	testdata := make(Environment)
	result, err := ReadDir(tmp)
	require.NoError(t, err)
	require.Equal(t, result, testdata)
}
