package write

import (
	"fmt"
	"os"
)

type WriteFileInit struct {
	data     []byte
	filename string
}

func WriteFile(writeFileInit WriteFileInit) error {
	err := os.WriteFile(
		writeFileInit.filename,
		writeFileInit.data,
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Printf("error writing to file: %v", err)
		return err
	}

	return nil
}
