package write

import (
	"fmt"
	"os"
)

func WriteFile(data []byte, filename string) error {
	err := os.WriteFile(filename, data, os.FileMode(0644))

	if err != nil {
		fmt.Printf("error writing to file: %v", err)
		return err
	}

	return nil
}
