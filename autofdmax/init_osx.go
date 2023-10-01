//go:build darwin

package autofdmax

import "github.com/khulnasoft-lab/fdmax"

func init() {
	_ = fdmax.Set(fdmax.OSXMax)
}
