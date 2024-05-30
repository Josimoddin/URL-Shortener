package models

import (
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
	"urlShortener/src/dao"
	"urlShortener/src/entities"
	"urlShortener/src/logger"
)

var lock = sync.Mutex{}

func ShortenUrl(req *entities.URLRequest) (string, bool, error, string) {
	logger.Log.Println("In side ShortenUrl")
	lock.Lock()
	defer lock.Unlock()
	shortURL := generateShortURL()
	domain := extractDomain(req.URL)

	shortURL = dao.SaveURL(shortURL, domain, req.URL)
	return shortURL, true, nil, "Url shortened successfully"
}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func extractDomain(url string) string {
	// Basic domain extraction logic
	// In a real-world application, use a proper URL parser
	parts := strings.Split(url, "/")
	if len(parts) > 2 {
		return parts[2]
	}
	return url
}

func GetOriginalURL(shortURL string) (string, bool) {
	orgUrl, exist := dao.GetOriginalURL(shortURL)
	if exist {
		return orgUrl, true
	}
	return "", false
}

func GetMetrics() []entities.MetricsResponse {
	topDomains := dao.GetTopDomains()

	var metrics []entities.MetricsResponse
	for domain, count := range topDomains {
		metrics = append(metrics, entities.MetricsResponse{Domain: domain, Count: count})
	}

	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Count > metrics[j].Count
	})

	if len(metrics) > 3 {
		metrics = metrics[:3]
	}
	return metrics
}
