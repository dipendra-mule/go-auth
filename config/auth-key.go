package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateRandomKey() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Failed to generate random key: ", err)
	}

	return base64.URLEncoding.EncodeToString(b)
}