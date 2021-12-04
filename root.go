package amusing_x

import (
	"path"
	"runtime"
)

func Root() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)
	return root
}
