package root

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// this global basepath makes it possible to find .env file in testing
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

// Get specific env variable from any package
func GetEnvVar(v string) (string, error) {
	basepath := RootDir()
	err := godotenv.Load(filepath.Join(basepath, "../.env"))
	if err != nil {
		return "", err
	}
	res := os.Getenv(v)
	return res, nil
}
