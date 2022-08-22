package testing_init

import (
	"os"
	"path"
	"runtime"
)

// hook for correct template path
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
