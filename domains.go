package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// readLines reads lines in a file
// code is from https://stackoverflow.com/a/29314318
func readLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}

	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}

	return lines, nil
}

// Domains returns a list of domains in the email_provider_domains.txt file
func Domains() []string {
	var results []string

	// read lines
	lines, err := readLines("email_provider_domains.txt")
	if err != nil {
		log.Fatal("cannot read a file with domains")
	}

	// validate lines
	for _, line := range lines {
		if !strings.Contains(line, ".") {
			continue
		}

		// remove the "\n" character
		line = strings.TrimSpace(line)

		results = append(results, line)
	}

	return results
}
