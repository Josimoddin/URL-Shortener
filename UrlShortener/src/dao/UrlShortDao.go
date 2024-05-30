package dao

type URLStore struct {
	urls   map[string]string
	counts map[string]int
}

var store = URLStore{
	urls:   make(map[string]string),
	counts: make(map[string]int),
}

func SaveURL(shortURL, domain, originalURL string) {

	store.urls[shortURL] = originalURL
	store.counts[domain]++
}

func (s *URLStore) GetOriginalURL(shortURL string) (string, bool) {
	url, exists := s.urls[shortURL]
	return url, exists
}
