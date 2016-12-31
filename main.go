package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Print an error message to stderr, and exit the program.
func errorExit(exitCode int, template string, elements ...interface{}) {
	fmt.Fprintf(os.Stderr, template+"\n", elements...)
	os.Exit(exitCode)
}

func main() {
	var leftFilePath, rightFilePath string
	var leftCol, rightCol int
	var trimAndIgnoreCase bool
	flag.StringVar(&leftFilePath, "left", "", "Path to the left file")
	flag.StringVar(&rightFilePath, "right", "", "Path to the right file")
	flag.IntVar(&leftCol, "leftcol", 0, "Index of column to join from left file. First column is 0.")
	flag.IntVar(&rightCol, "rightcol", 0, "Index of column to join from right file. First column is 0.")
	flag.BoolVar(&trimAndIgnoreCase, "trimandignorecase", false, "Perform case insensitive join, and ignore prefix/suffix spaces.")
	flag.Parse()
	if leftFilePath == "" || rightFilePath == "" {
		fmt.Println("csvjoin: join two CSV files and print result to standard output")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	leftFile, err := os.Open(leftFilePath)
	if err != nil {
		errorExit(1, "Left file error: %v", err)
		return
	}
	rightFile, err := os.Open(rightFilePath)
	if err != nil {
		errorExit(1, "Right file error: %v", err)
		return
	}

	leftReader := csv.NewReader(leftFile)
	rightReader := csv.NewReader(rightFile)
	if trimAndIgnoreCase {
		leftReader.LazyQuotes = true
		leftReader.TrimLeadingSpace = true
		rightReader.LazyQuotes = true
		rightReader.TrimLeadingSpace = true
	}
	left, err := leftReader.ReadAll()
	if err != nil {
		errorExit(1, "Left file error: %v", err)
		return
	}
	right, err := rightReader.ReadAll()
	if err != nil {
		errorExit(1, "Right file error: %v", err)
		return
	}

	result := Join(left, right, leftCol, rightCol, trimAndIgnoreCase)
	writer := csv.NewWriter(os.Stdout)
	writer.WriteAll(result)
	writer.Flush()
}
