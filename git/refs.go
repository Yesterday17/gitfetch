package git

import (
	"github.com/Yesterday17/gitfetch/utils"
)

func (f *fetcher) FetchRef(path string) (hash string) {
	hash = utils.GetText(f.base + path)

	if len(hash) != 40 {
		hash = ""
	}
	return
}
