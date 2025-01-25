package write

import (
	"fmt"
	"os"
)

// WriteFile writes the provided data to a file with the specified filename.
// It sets the file permissions to 0644 by default. If an error occurs during
// the write operation, it prints an error message and returns the error.
//
// Parameters:
//   - data: The byte slice containing the data to be written to the file.
//   - filename: The name of the file to which the data will be written.
//
// Returns:
//   - error: An error if the write operation fails, otherwise nil.
func WriteFile(data []byte, filename string) error {
	err := os.WriteFile(filename, data, os.FileMode(0644))

	if err != nil {
		fmt.Printf("error writing to file: %v", err)
		return err
	}

	return nil
}
