package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Print("too few arguments")
		return
	}

	env, err := ReadDir(args[1])
	if err != nil {
		log.Printf("error reading envs %s", err.Error())
	}

	result := RunCmd(args[2:], env)
	os.Exit(result)
}
