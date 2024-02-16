//go:build wasm

package cage

func (rl *ResLoader) loadBytes(url string) ([]byte, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	defer func() { _ = resp.Body.Close() }()
	return io.ReadAll(resp.Body)
}
