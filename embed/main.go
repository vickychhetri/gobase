package main

import (
	_ "embed" // Required for using //go:embed directive
	"fmt"
	"os"
)

/*
	Using go:embed to include files into the compiled binary.

	IMPORTANT:
	- go:embed must be placed directly above a variable declaration.
	- The variable type decides how the file will be stored:
		[]byte → raw bytes (good for images, PDFs, binary files)
		string → text content (for .txt, .json, .html)
*/

// Embed a PNG image into the binary as raw bytes.
//
//go:embed h.png
var himage []byte

// Embed text content as a normal string.
//
//go:embed infome.txt
var infome string

// Embed this main.go file itself as a string (good for self-inspection or debugging).
//
//go:embed main.go
var embedcode string

func main() {

	// Inform the user that the embedded content will be written somewhere.
	fmt.Println("File kept in embed")

	// Write the embedded PNG image to a new file.
	// We ignore the returned integer using "_" because we don't need it here.
	_ = WriteFile("temp_of_h.png", himage)

	// Print the text content that was embedded from infome.txt
	fmt.Println(infome)

	// Print the content of main.go which we embedded.
	fmt.Println(embedcode)
}

/*
WriteFile writes raw byte data (img []byte) to a file at the given path.

PARAMETERS:
- path: file path where data will be written
- img:  raw binary data (image in this case)

RETURNS:
- number of bytes written to the file
*/
func WriteFile(path string, img []byte) int {

	/*
		Open or create the file using os.OpenFile.

		os.O_CREATE → create file if it does NOT exist.
		os.O_WRONLY → open file for WRITE ONLY.
		644         → file permission in octal. Equivalent to:
		              - Owner: read + write
		              - Group: read
		              - Others: read
	*/
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// If file cannot be created or opened, stop the program.
		panic(err)
	}

	// Write the byte array to the file.
	n, err := fd.Write(img)
	if err != nil {
		panic(err)
	}

	// Print how many bytes were written.
	fmt.Println("Size added to file: ", fmt.Sprintf("%d", n))

	return n
}
