package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // Пример использования клиента
    originalURL := "https://google.com"
    shortURL, err := shortenURL(originalURL)
    if err != nil {
        fmt.Println("Ошибка при сокращении URL:", err)
        return
    }

    fmt.Printf("Сокращенный URL: http://localhost:8080/%s\n", shortURL)

    // Перенаправление на оригинальный URL
    err = redirectURL(shortURL)
    if err != nil {
        fmt.Println("Ошибка при перенаправлении:", err)
        return
    }
}

// Функция для сокращения URL
func shortenURL(originalURL string) (string, error) {
    requestBody, err := json.Marshal(map[string]string{
        "url": originalURL,
    })
    if err != nil {
        return "", err
    }

    resp, err := http.Post("http://localhost:8080/shorten", "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var result map[string]string
    if err := json.Unmarshal(body, &result); err != nil {
        return "", err
    }

    return result["short_url"], nil
}

// Функция для перенаправления на оригинальный URL
func redirectURL(shortURL string) error {
    resp, err := http.Get("http://localhost:8080/" + shortURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusMovedPermanently {
        return fmt.Errorf("не удалось перенаправить на оригинальный URL")
    }

    fmt.Println("Перенаправление на:", resp.Header.Get("Location"))
    return nil
}