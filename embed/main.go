package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed h.png
var himage []byte

//go:embed infome.txt
var infome string

//go:embed main.go
var embedcode string

func main() {
	fmt.Println("File Keep in embed")
	_ = WriteFile("temp_of_h.png", himage)

	fmt.Println(infome)
	fmt.Println(embedcode)
}

func WriteFile(path string, img []byte) int {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		panic(err)
	}

	n, err := fd.Write(img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Size Added in file: ", fmt.Sprintf("%d", n))
	return n
}
