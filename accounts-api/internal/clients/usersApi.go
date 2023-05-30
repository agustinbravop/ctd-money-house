package clients

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

var (
	ErrNotValid = errors.New("token is not valid")
)

// ValidateUserToken manda el tokenHeader al endpoint /users/validate y devuelve las claims del JWT.
// Si el tokenHeader es un JWT inválido, ValidateUserToken devuelve un error.
func ValidateUserToken(tokenHeader string) (map[string]string, error) {
	host := os.Getenv("USERS_API_HOST")
	req, err := http.NewRequest(http.MethodPost, host+"/api/v1/auth/validate", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", tokenHeader)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		return nil, ErrNotValid
	}
	var claims map[string]string
	err = json.NewDecoder(res.Body).Decode(&claims)
	return claims, nil
}
