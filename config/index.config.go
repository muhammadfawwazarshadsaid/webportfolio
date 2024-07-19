package config

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	// StaticDir is the directory where static files are stored
	StaticDir string
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	BasePath := filepath.Dir(b)
	StaticDir = filepath.Join(BasePath, "../public")

	// Create the static directory if it doesn't exist
	if _, err := os.Stat(StaticDir); os.IsNotExist(err) {
		os.MkdirAll(StaticDir, os.ModePerm)
	}
}
