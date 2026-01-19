package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// set flag for multiplelines such as -multiple --m and file
	help := flag.Bool("help", false, "Display help")
	multiline := flag.Bool("multiline", false, "Enable multi-line decoding mode")
	encodeFlag := flag.Bool("encode", false, "Encode partten")
	file := flag.String("file", "", "Decode input file")
	// parse all flags
	flag.Parse()

	if *help {
		displayHelp()
		return
	}
	// check for multiline mode and handle flags
	if *multiline {
		if *encodeFlag {
			encodeMultipleLines() // encode multiple lines of partten
		} else {
			// read and decode multiline mode from stdin
			decodeMultipleLines()
		}
		return
	}
	// check for file flag
	if *file != "" {
		if _, err := os.Stat(*file); os.IsNotExist(err) {
			fmt.Println("Error \nFile does not exist:", err)
			os.Exit(1)
		}
		if *encodeFlag {
			encodeFile(*file)

		} else {
			decodeFile(*file)
		}
		return
	}

	arg := flag.Args()
	// handle flags
	if len(arg) < 1 {
		displayHelp()
		os.Exit(1)
	}
	// get input from arg 0
	input := arg[0]

	if *encodeFlag {
		result, err := encoder(input)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		fmt.Println(result) // print the result of encoding
	} else {
		result, err := decoder(input)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		fmt.Println(result) // print the result of decoding

	}

}

// func help flag
func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run . '[5 $]'               → Decode single input")
	fmt.Println("  go run . -encode 'argument'     → Encode single input")
	fmt.Println("  go run . -multiline             → Multi-line decode (add done)")
	fmt.Println("  go run . -encode -multiline    → Multi-line encode")
	fmt.Println("  go run . -file=input.txt        → Decode file")
	fmt.Println("  go run . -encode -file=input.txt → Encode file")
}
