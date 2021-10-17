package config

import (
    "net/url"
    "os"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func GetBackendUrl() *url.URL {
    backendEnv := getEnv("BACKEND_URL", "http://localhost:8080")
    backendUrl, err := url.Parse(backendEnv)
    if err != nil {
        panic(err)
    }
    return backendUrl
}
    
func GetBackendType() string {
    return getEnv("BACKEND_TYPE", "vespa")
}

func GetListenAddress() string {
    port := getEnv("PORT", "80")
    return ":" + port
}

