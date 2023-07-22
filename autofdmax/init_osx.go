//go:build darwin

package autofdmax

import "github.com/khulnasoft-labs/fdmax"

func init() {
	_ = fdmax.Set(fdmax.OSXMax)
}
