package write

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	filename := "test-file.txt"
	testdata := []byte("test data")

	err := WriteFile(testdata, filename)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	writtenTestData, testErr := os.ReadFile(filename)
	if testErr != nil {
		t.Errorf("error reading test data: %v", testErr)
	}

	if string(writtenTestData) != "test data" {
		t.Errorf("expected 'test data', got %s", string(writtenTestData))
	}

	os.Remove(filename)
}
