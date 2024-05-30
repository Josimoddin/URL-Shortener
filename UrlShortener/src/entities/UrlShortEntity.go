package entities

import (
	"encoding/json"
	"io"
)

type URLRequest struct {
	URL string `json:"url"`
}

// type URLResponse struct {
// 	ShortURL string `json:"short_url"`
// }

type URLResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	ShortUrl string `json:"shorturl"`
}
type MetricsResponse struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

func (w *URLRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
