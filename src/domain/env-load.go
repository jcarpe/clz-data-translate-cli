package domain

import (
	"os"
	"strings"
)

// LoadEnv reads env variables from a file and sets them in the current process.
// The file should contain key-value pairs in the format KEY=VALUE, one per line.
// Quoted values will have their quotes removed.
//
// Parameters:
//   - filePath: The path to the file containing the environment variables.
//
// Panics:
//   - If there is an error reading the file, the function will panic.
func LoadEnv(filePath string) {
	envData, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	data := string(envData)
	dataPairs := strings.Split(data, "\n")

	for i, pair := range dataPairs {
		dataPairs[i] = strings.TrimSpace(pair)
		keyVal := strings.Split(pair, "=")
		val := strings.TrimFunc(keyVal[1], func(r rune) bool {
			return r == '"'
		})

		os.Setenv(keyVal[0], val)
	}
}
