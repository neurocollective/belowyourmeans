package cookie

import (
	"crypto/rand"
	"encoding/hex"
	"neurocollective.io/neurocollective/belowyourmeans/server/constants"
	"os"
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
		return constants.COOKIE_KEY + "=" + cookie
	}

	return constants.COOKIE_KEY + "=" + cookie + "; HttpOnly; Max-Age=3600; SameSite=Strict; Secure"
}
