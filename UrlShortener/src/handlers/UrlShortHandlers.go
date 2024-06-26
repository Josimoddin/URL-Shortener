package handlers

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"strings"
	"urlShortener/src/entities"
	"urlShortener/src/logger"
	"urlShortener/src/models"
)

func ThrowShortenUrlResponse(successMessage string, responseData string, w http.ResponseWriter, success bool) {
	var response = entities.URLResponse{}
	response.Success = success
	response.Message = successMessage
	response.ShortUrl = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ShortenUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.URLRequest{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.ShortenUrl(&data)
		ThrowShortenUrlResponse(msg, data, w, success)
	}
}

func RedirectUrl(w http.ResponseWriter, req *http.Request) {
	logger.Log.Println("Fetching MstAutoForward Data")
	var query = req.URL.Path
	// fmt.Println("query", query)
	s := strings.Split(query, "/")
	// fmt.Println("split", s, len(s))
	url := ""
	if len(s) == 4 {
		url = s[3]
	}
	// fmt.Println("URLLLL", url)
	originalURL, exists := models.GetOriginalURL(url)
	if !exists {
		http.NotFound(w, req)
		return
	}
	http.Redirect(w, req, originalURL, http.StatusMovedPermanently)
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {

	metrics := models.GetMetrics()

	json.NewEncoder(w).Encode(metrics)
}
