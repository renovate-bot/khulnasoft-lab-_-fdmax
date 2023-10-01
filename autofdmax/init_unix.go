//go:build linux || openbsd || netbsd

package autofdmax

import (
	"github.com/khulnasoft-lab/fdmax"
)

func init() {
	_ = fdmax.Set(fdmax.UnixMax)
}
