package password

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io"

	"neurocollective.io/neurocollective/belowyourmeans/server/structs"
)

func HashPassword(password string) (string, error) {

	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GetSignupPayload(c *gin.Context) (*structs.SignupPayload, error) {

	jsonBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	payload := new(structs.SignupPayload)

	err = json.Unmarshal(jsonBytes, payload)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
