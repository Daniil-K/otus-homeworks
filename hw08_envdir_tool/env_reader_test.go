package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadTestDir(t *testing.T) {
	testdata := make(Environment)
	testdata["BAR"] = EnvValue{"bar", false}
	testdata["EMPTY"] = EnvValue{"", false}
	testdata["FOO"] = EnvValue{"   foo\nwith new line", false}
	testdata["HELLO"] = EnvValue{"\"hello\"", false}
	testdata["UNSET"] = EnvValue{"", true}
	result, err := ReadDir("testdata/env")
	fmt.Println(result)
	fmt.Println(testdata)
	require.Nil(t, err)
	require.Equal(t, result, testdata)
}

func TestEmptyDir(t *testing.T) {
	tmp, err := os.MkdirTemp("", "temp")
	require.NoError(t, err)

	defer func() {
		err := os.Remove(tmp)
		if err != nil {
			log.Fatal(err)
		}
	}()
	testdata := make(Environment)
	result, err := ReadDir(tmp)
	require.Nil(t, err)
	require.Equal(t, result, testdata)
}
