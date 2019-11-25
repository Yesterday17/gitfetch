package git

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/Yesterday17/gitfetch/utils"
	"io"
)

func (f *fetcher) GenerateObjectPath(hash string) string {
	return fmt.Sprintf("%sobjects/%s/%s", f.base, hash[0:2], hash[2:])
}

func (f *fetcher) FetchObject(hash string) []byte {
	bin := utils.GetBinary(f.GenerateObjectPath(hash))
	r, err := zlib.NewReader(bytes.NewReader(bin))
	if err != nil {
		return nil
	}

	defer r.Close()
	buffer := bytes.Buffer{}
	_, _ = io.Copy(&buffer, r)
	return buffer.Bytes()
}
