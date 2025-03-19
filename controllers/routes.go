package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Wasay1567/url-shortner-golang/models"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	var payload models.GetUrl
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	s := generateShortUrl(payload.Url)

	u := models.URL{
		OrginalUrl: payload.Url,
		ShortenUrl: s,
		CreatedAt:  time.Now(),
	}

	db.Create(&u)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"short_url": u.ShortenUrl})
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	var u models.URL
	err := db.Where("shorten_url = ?", code).First(&u).Error
	if err != nil {
		http.Error(w, "Invalid code", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, u.OrginalUrl, http.StatusFound)

}

func generateShortUrl(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash
}
