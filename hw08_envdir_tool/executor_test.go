package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		code  int
	}{
		{"empty call", []string{""}, 0},
		{"empty echo", []string{"echo"}, 0},
		{"command not found", []string{"ecco"}, 127},
		{"sh", []string{"./testdata/echo.sh"}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code := RunCmd(test.input, nil)
			require.Equal(t, code, test.code)
		})
	}
}

func TestRunCmdErr2(t *testing.T) {
	err := os.WriteFile("tmp.sh", []byte("#!/bin/bash\nexit 2\n"), 0777) //nolint:gofumpt
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := os.Remove("tmp.sh")
		require.NoError(t, err)
	}()

	code := RunCmd([]string{"./tmp.sh"}, nil)
	require.Equal(t, code, 2)
}

func TestRunMvCmd(t *testing.T) {
	tmp, err := os.MkdirTemp("", "temp")
	require.NoError(t, err)

	defer func() {
		err := os.RemoveAll(tmp)
		require.NoError(t, err)
	}()

	nameBefore := filepath.Join(tmp, "tmp")
	nameAfter := filepath.Join(tmp, "tmp_renamed")
	err = os.WriteFile(nameBefore, []byte("just some file"), 0777) //nolint:gofumpt
	require.NoError(t, err)

	code := RunCmd([]string{"mv", nameBefore, nameAfter}, nil)

	contents, err := os.ReadDir(tmp)
	require.NoError(t, err)

	require.Equal(t, code, 0)
	require.Len(t, contents, 1)
	require.Equal(t, contents[0].Name(), "tmp_renamed")
}

func TestWithEnvRunCmd(t *testing.T) {
	err := os.WriteFile("tmp.sh", []byte("#!/bin/bash\necho $FOO > tmp.out"), 0777) //nolint:gofumpt
	require.NoError(t, err)
	defer func() {
		err := os.Remove("tmp.sh")
		require.NoError(t, err)
	}()

	env := make(Environment)
	env["FOO"] = EnvValue{"text", false}
	code := RunCmd([]string{"./tmp.sh"}, env)
	text, err := os.ReadFile("tmp.out")
	require.NoError(t, err)

	defer func() {
		err := os.Remove("tmp.out")
		require.NoError(t, err)
	}()

	require.Equal(t, string(text), "text\n")
	require.Equal(t, code, 0)
}
