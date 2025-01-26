package domain

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvLoad(t *testing.T) {
	// Setup
	filePath := "./_test/.env.test"

	// Execution
	LoadEnv(filePath)

	fmt.Println(os.Getenv("TEST_VALUE"))

	// Validation
	if os.Getenv("TEST_VALUE") != "test-value" {
		t.Errorf("Expected %s but got %s", "test-value", os.Getenv("TEST_VALUE"))
	}

	// Cleanup
	os.Unsetenv("TEST_VALUE")
}
