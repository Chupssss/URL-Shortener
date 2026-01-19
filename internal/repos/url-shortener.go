package repos

type UrlRepo struct {
	data map[string]string
}

func NewUrlRepos() *UrlRepo {
	return &UrlRepo{
		data: make(map[string]string),
	}
}

func (r *UrlRepo) Save(shortUrl, originalUrl string) {
	r.data[shortUrl] = originalUrl
}

func (r *UrlRepo) Get(shortUrl string) (string, bool) {
	url, ok := r.data[shortUrl]
	return url, ok
}
