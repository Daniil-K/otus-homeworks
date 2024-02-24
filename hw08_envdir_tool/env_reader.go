package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, file := range files {
		fi, err := os.Stat(filepath.Join(dir, file.Name()))
		if err != nil {
			log.Printf("can`t get file info \"%s\", %s", file.Name(), err.Error())
			continue
		}

		if !fi.Mode().IsRegular() {
			continue
		}

		name := file.Name()
		if strings.Contains(name, "=") {
			log.Printf("filename \"%s\" contains \"=\"", name)
			continue
		}

		value, err := readEnv(filepath.Join(dir, name))
		if err != nil {
			log.Printf("can`t read file \"%s\", %s", name, err.Error())
			continue
		}

		env[name] = *value
	}

	return env, nil
}

func readEnv(file string) (*EnvValue, error) {
	valueBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if len(valueBytes) == 0 {
		return &EnvValue{"", true}, nil
	}

	str := string(valueBytes)
	if strings.Contains(str, "\n") {
		str = strings.Split(str, "\n")[0]
	}
	str = strings.ReplaceAll(str, "\x00", "\n")
	value := strings.TrimRight(strings.TrimRight(str, " "), "\t")
	env := EnvValue{
		Value:      value,
		NeedRemove: false,
	}
	return &env, nil
}
