package write

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	// Arrange
	writeInit := WriteFileInit{
		data:     []byte("test data"),
		filename: "test-file.txt",
	}

	// Act
	err := WriteFile(writeInit)

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	writtenTestData, testErr := os.ReadFile(writeInit.filename)
	if testErr != nil {
		t.Errorf("error reading test data: %v", testErr)
	}

	if string(writtenTestData) != "test data" {
		t.Errorf("expected 'test data', got %s", string(writtenTestData))
	}

	// Clean up
	// os.Remove(writeInit.filename)
}
