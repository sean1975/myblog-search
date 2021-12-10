package utils

import (
	"testing"
)

func TestGzip(t *testing.T) {
	plainString := "Hello world"
	plainBytes := []byte(plainString)
	compressed := Gzip(plainBytes)
	decompressed := Gunzip(compressed)
	newString := string(decompressed)
	if plainString != newString {
		t.Errorf("Failed to compress or decompress")
	}
}
