package storage

import (
	"fmt"
    "math/rand"
    "strings"
    "sync"
)

var (
    urls = make(map[string]string)
    mu   sync.RWMutex
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortURL() string {
    var sb strings.Builder
    for i := 0; i < 6; i++ {
        sb.WriteByte(charset[rand.Intn(len(charset))])
    }
    return sb.String()
}

func SaveURL(originalURL string) string {
    mu.Lock()
    defer mu.Unlock()

    shortURL := generateShortURL()
    urls[shortURL] = originalURL
    return shortURL
}

func GetURL(shortURL string) (string, bool) {
    mu.RLock()
    defer mu.RUnlock()

    fmt.Printf("Current URLs in storage: %v\n", urls) 
    originalURL, exists := urls[shortURL]
    return originalURL, exists
}