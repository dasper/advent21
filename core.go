package advent21

import (
	"fmt"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	return file, err
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file properly:", err.Error())
	}
}