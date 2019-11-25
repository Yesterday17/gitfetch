package git

import (
	"github.com/Yesterday17/gitfetch/utils"
	"regexp"
)

// FetchHead returns ref name of current head
func (f *fetcher) FetchHead() (head, raw string) {
	raw = utils.GetText(f.base + HeadPath)
	r := regexp.MustCompile(HeadRegexp)

	match := r.FindStringSubmatch(raw)
	if len(match) != 2 {
		raw = ""
	} else {
		head = match[1]
	}
	return
}

func (f *fetcher) FetchOriginHead() (hash string) {
	return f.FetchRef(OriginHeadPath)
}
