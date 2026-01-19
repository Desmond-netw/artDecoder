# ASCII ArtDecoder 

## A command line tool that converts art data into text-based art.
A simple command-line tool that **encodes** and **decodes** custom pattern expressions such as:
```bash
[5 $] → $$$$$
[10 o] → oooooooooo
```
## The situation
A local artist named Chris has asked you to help them create a project to simplify generating text-based images.Chris said Very often, I have an idea for an image, but it requires so much repetition. I wish there was a quicker way to create those images.

## Supports:
- Single-line decoding
- Single-line encoding
- Multi-line mode (stdin streaming)
- File input mode
- Full error handling for malformed patterns

## Features
- Encode pattern → `[count text]`
- Decode pattern → expands repetitions
- Multi-line encode/decode
- File input: process each line in a file
- Detects malformed input:
  - missing space separator
  - non-numeric repeat count
  - empty repeat string
  - nested brackets
  - **Unbalanced or stray brackets**

## Rules:

* Count must be a number
* There must be exactly one space between count and string
* Repeat string may contain any characters, including spaces
* Nested brackets like [3 [x]] are not allowed
* It displays "Error" if square brackets marks are unbalanced
* On any error → prints:

## Usage

### **Decode a single input**
```sh
go run . '[5 $]'
```
### **Encode a single input**
```sh
go run . -encode 'hello'
```
### **Multi-line decode (stdin stream)**
NOTE:: Always end multi-line with  `done`
```sh
go run . -multiline

[7 %]
[4 ^] 
[7 %]
done
%%%%%%%
^^^^
%%%%%%%
```

### **Multi-line encode**
To toggle encoding mode `-encode` before `-multiline`
```sh
go run . -encode -multiline
```
### **Decode from File**
```sh
go run . -file input.txt
```
### **Encode file contents**
```sh
go run . -encode -file input.txt
```

## Help Command
 Use 
 ```sh
 go run . -help
 ```

 ### Examples

 ```vbnet
Input:  "[3 o]"
Output: "ooo"

Input:  "[10  *]"
Output: "**********"

Input: "abc[2 d]xyz"
Output: "abcddxyz"
``` 