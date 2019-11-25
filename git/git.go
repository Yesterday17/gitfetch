package git

type Fetcher interface {
	FetchHead() (head, raw string)
	FetchOriginHead() (hash string)
	FetchRef(path string) (hash string)
	FetchObject(hash string) (out []byte)
}

type fetcher struct {
	base string
}

func NewFetcher(base string) Fetcher {
	return &fetcher{
		base,
	}
}
