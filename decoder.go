package main

// importing package
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// decoderPattern
func decoder(input string) (string, error) {

	/* ------------- Global check fo unbalance brackets ----- */
	openCount := 0
	closeCount := 0

	for _, ch := range input {
		if ch == '[' {
			openCount++
		} else if ch == ']' {
			closeCount++

			// handle closing bracket appearing before pning bracket
			if closeCount > openCount {
				return "", fmt.Errorf("Error")
			}
		}
	}

	// handle more '[' and ']'
	if openCount != closeCount {
		return "", fmt.Errorf("Error")
	}

	//---------  var to build the final decoded string
	var resultBuilder strings.Builder // strings concatenating
	currentIndex := 0
	inputLength := len(input)

	//---------itereate over input lenght as far as currentIndex is less than inputLenght
	for currentIndex < inputLength {
		// find nex open bracket '[' from the current position
		openBracketPos := strings.IndexByte(input[currentIndex:], '[')
		if openBracketPos == -1 {

			resultBuilder.WriteString(input[currentIndex:]) //no more bracket append the rest
			break                                           // end the loop -> no '['
		}

		//append any char that appear before '['
		anyTextStart := currentIndex
		anyTextEnd := currentIndex + openBracketPos
		if anyTextEnd > anyTextStart {
			resultBuilder.WriteString(input[anyTextStart:anyTextEnd])
		}

		// get the actual index of strings after '['
		contentStartPos := anyTextEnd + 1

		// find the corresponding closeBrackt ']'
		closeBracketPos := strings.IndexByte(input[contentStartPos:], ']')
		if closeBracketPos == -1 {
			return "", fmt.Errorf("unbalanced square brackets")
		}

		// get the actual index of string before ']'
		contentEndPos := contentStartPos + closeBracketPos

		// extract the actual content between  '[' and ']'
		mainContent := input[contentStartPos:contentEndPos]

		// disallow nested brackets
		if strings.ContainsAny(mainContent, "[]") {
			return "", fmt.Errorf("nested bracket not allowed")
		}

		// check for space between mainContent
		spaceIndex := strings.IndexByte(mainContent, ' ')
		if spaceIndex == -1 {
			return "", fmt.Errorf("missing space separator")
		}
		// split maincontent to counter and char to repeat
		countStr := strings.TrimSpace(mainContent[:spaceIndex])
		// repeatStr := strings.TrimSpace(mainContent[spaceIndex+1:])
		repeatStr := mainContent[spaceIndex+1:]
		// validate countStr and repeatStr
		if countStr == "" {
			return "", fmt.Errorf("missing repeat count")
		}
		if len(repeatStr) == 0 {
			return "", fmt.Errorf("missing string to repeat")
		}
		// convert count from countstr
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return "", fmt.Errorf("repeat count is not a number")
		}
		if count < 0 {
			return "", fmt.Errorf("negative repeat count")
		}

		//Append the repeated char to the final result
		for j := 0; j < count; j++ {
			resultBuilder.WriteString(repeatStr)
		}

		// move the current index to pass the current index
		currentIndex = contentEndPos + 1
	}

	return resultBuilder.String(), nil
}

// decode multiple lines
// read multiple line from stdin and decodes each separately
func decodeMultipleLines() {
	result := []string{}
	// initilize scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // skip empty line
		}
		if line == "done" {
			break
		}
		// use exit to close application sefely
		if line == "exit" || line == "EXIT" || line == "end" {
			os.Exit(0)
		}
		output, err := decoder(line)
		result = append(result, output)
		if err != nil {
			fmt.Println("error")
			continue
		}

	}
	fmt.Println(strings.Join(result, "\n"))
	if err := scanner.Err(); err != nil {
		fmt.Println("error")
	}
}

// decode file mode
// read multiple line from file and decodes each line
func decodeFile(filename string) {

	// open the file first
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(1)
	}
	defer file.Close()

	// use bufio scanner to scan and read the file line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileResult, err := decoder(line)
		if err != nil {
			fmt.Println("error")
			return
		}
		fmt.Println(fileResult)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file")
	}
}
