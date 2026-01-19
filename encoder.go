package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Encoding will put the code into orginal format
// Step:
/*
  1. This takes each character in a runes
  2. check if that char repeat
*/

// encoding the decodepartten into the form [n char] and leave the rest
func encoder(input string) (string, error) {
	if len(input) == 1 {
		return "", fmt.Errorf("empty input")
	}
	var mybuilder strings.Builder // to concatenate strings

	// use rune to hold all chars
	runes := []rune(input)
	lengthRunes := len(runes)

	count := 1
	// iterate through input runes
	for index := 0; index < lengthRunes; index++ {
		// start counting repeats
		if index+1 < lengthRunes && runes[index] == runes[index+1] {
			count++
			continue
		}

		// if char repeated more than once, encode it
		// else append it as it is
		if count > 1 {
			mybuilder.WriteString(fmt.Sprintf("[%d %c]", count, runes[index]))
		} else {
			mybuilder.WriteString(string(runes[index]))
		}

		count = 1 // reset count for next rune
	}
	//return everthing as string
	return mybuilder.String(), nil
}

// encode multiple lines reading line by line
func encodeMultipleLines() {
	result := []string{}
	// declare scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// eachline in stdin
		eachline := scanner.Text()

		if eachline == "" {
			continue
		}
		if eachline == "done" {
			break
		}
		// exiting with done or exit
		if strings.ToLower(eachline) == "exit" {
			os.Exit(1)
		}

		output, err := encoder(eachline)
		result = append(result, output)
		if err != nil {
			fmt.Println("error")
			continue
		}

	}
	fmt.Println(strings.Join(result, "\n"))
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input", err)
	}
}

// encode file parttens
func encodeFile(filename string) {
	// open the file name
	myfile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		os.IsNotExist(err)
		os.Exit(1)
	}
	defer myfile.Close() // ensure file close

	// usin bufio to scan the file content
	scanner := bufio.NewScanner(myfile)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		encoded, err := encoder(line)
		if err != nil {
			fmt.Println("Error encoding")
			return
		}
		fmt.Println(encoded) // display encoded chars in file
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
	}
}
