package handlers

import (
    "encoding/json"
    "net/http"
    "url-shorter/storage"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req struct {
        URL string `json:"url"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    shortURL := storage.SaveURL(req.URL)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[1:] 
    originalURL, exists := storage.GetURL(shortURL)
    if !exists {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}