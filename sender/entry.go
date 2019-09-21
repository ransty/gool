package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var fileToSend string // abs returns a string of the full file path

func main() {
	const (
		usage = "The file to send"
	)
	// Take a string input of a file name
	flag.StringVar(&fileToSend, "file", "", usage)
	flag.StringVar(&fileToSend, "f", "", usage+" (shorthand)")
	flag.Parse()

	abs := getFilePath(fileToSend)

	fmt.Println("file:", abs)
}

// Get the file path of the input
// This function should return an absolute path
// to a file that exists on the fs
// Will error out if the input doesn't exist
// or if the input is a directory
func getFilePath(input string) string {
	abs, err := filepath.Abs(input)
	if err != nil {
		log.Fatal(err)
	}
	if stat, err := os.Stat(abs); os.IsNotExist(err) {
		log.Fatal(abs, " does not exist on the file system")
	} else if stat.IsDir() {
		log.Fatal(abs, " is a directory, not a file")
	}
	return abs
}
