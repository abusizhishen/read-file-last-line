package main

import (
	"fmt"
	"log"
	"os"

	read_file_last_line "github.com/abusizhishen/read-file-last-line"
)

func main() {
	if len(os.Args) <= 1 {
		println("no file specified")
		os.Exit(0)
	}

	file := os.Args[1]
	byt, n, err := read_file_last_line.ReadLastLine(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	log.Printf("offset: %d", n)

	fmt.Println(string(byt))
}
