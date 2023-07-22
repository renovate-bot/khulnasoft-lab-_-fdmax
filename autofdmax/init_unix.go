//go:build linux || openbsd || netbsd

package autofdmax

import (
	"github.com/khulnasoft-labs/fdmax"
)

func init() {
	_ = fdmax.Set(fdmax.UnixMax)
}
