package dao

import "fmt"

type URLStore struct {
	urls    map[string]string
	reverse map[string]string
	counts  map[string]int
}

var store = URLStore{
	urls:    make(map[string]string),
	reverse: make(map[string]string),
	counts:  make(map[string]int),
}

func SaveURL(shortURL, domain, originalURL string) string {
	url, exist := store.reverse[originalURL]
	if !exist {

		store.urls[shortURL] = originalURL
		store.reverse[originalURL] = shortURL
		store.counts[domain]++
		return shortURL
	}
	store.counts[domain]++
	fmt.Println(store)
	return url
}

func GetOriginalURL(shortURL string) (string, bool) {
	fmt.Println("store", store, shortURL)
	url, exists := store.urls[shortURL]
	return url, exists
}

func GetTopDomains() map[string]int {
	return store.counts
}
