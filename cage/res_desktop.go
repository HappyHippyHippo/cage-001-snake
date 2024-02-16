//go:build !wasm

package cage

import (
	"os"
)

func (rl *ResLoader) loadBytes(url string) ([]byte, error) {
	return os.ReadFile(url)
}
