package helper

import (
	"os"
	"path/filepath"
)

func PublicPath(paths ...string) string {
	cwd, _ := os.Getwd()
	all := append([]string{cwd, "public"}, paths...)
	return filepath.Join(all...)
}
