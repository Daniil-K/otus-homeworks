package main

import (
	"flag"
	"log"
	"strings"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	if len(strings.TrimSpace(from)) == 0 {
		log.Fatal("empty -from arg")
	}
	if len(strings.TrimSpace(to)) == 0 {
		log.Fatal("empty -to arg")
	}

	err := Copy(from, to, limit, offset)
	if err != nil {
		log.Fatal(err)
	}
}
