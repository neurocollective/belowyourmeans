package main

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"neurocollective.io/neurocollective/belowyourmeans/src/constants"
)

func GenerateCookie() (string, error) {

	size := 20
	byteSlice := make([]byte, size)
	_, err := rand.Read(byteSlice)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(byteSlice), nil
}

func GetSetCookieHeaderValue(cookie string) string {

	environment := os.Getenv("ENVIRONMENT")

	if environment == "dev" {
		return constants.COOKIE_KEY + "=" + cookie + "; HttpOnly; Max-Age=3600;"		
	}

	return constants.COOKIE_KEY + "=" + cookie + "; HttpOnly; Max-Age=3600; SameSite=Strict; Secure"
}
