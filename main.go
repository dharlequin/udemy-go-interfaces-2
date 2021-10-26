package main

import (
	"fmt"
	"io"
	"os"
)

type filePrinter struct{}

func main() {
	// skipping the first argument that is the programm itself
	fileNames := os.Args[1:]

	fp := filePrinter{}

	for i, fileName := range fileNames {
		file, err := os.Open(fileName)
		handleError(err, i)

		fmt.Printf("File #%v: \n", i+1)
		io.Copy(fp, file)
	}
}

func (fp filePrinter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs) + "\n")

	return 0, nil
}

func handleError(err error, index int) {
	if err != nil {
		fmt.Println(fmt.Sprintf("Error in file #%d:", index+1), err)
		os.Exit(1)
	}
}
