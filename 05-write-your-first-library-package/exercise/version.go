package version

import (
	"runtime"
)

func Version() string {
	return runtime.Version()
}
