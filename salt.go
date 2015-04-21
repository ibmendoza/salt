package salt

import (
	"encoding/json"
	"errors"
	"github.com/ibmendoza/cryptohelper"
	"time"
)

func GenerateKey() (string, error) {
	return cryptohelper.RandomKey()
}

func ExpiresInSeconds(d time.Duration) int64 {
	return time.Now().Add(time.Second * d).Unix()
}

func ExpiresInMinutes(d time.Duration) int64 {
	return time.Now().Add(time.Minute * d).Unix()
}

func ExpiresInHours(d time.Duration) int64 {
	return time.Now().Add(time.Hour * d).Unix()
}

//returns the corresponding claims as map[string]interface{} if token is valid
func Verify(token, naclKey string) (map[string]interface{}, error) {

	claims, err := cryptohelper.SecretboxDecrypt(token, naclKey)

	if err != nil {
		return nil, errors.New("Error decrypting claims")
	}

	mapClaims := make(map[string]interface{})

	if naclKey != "" {
		if err := json.Unmarshal([]byte(claims), &mapClaims); err != nil {

			return nil, err
		}
	}

	//if exp is not set, expiry is 0 (meaning no expiry)
	expiry, ok := mapClaims["exp"].(float64)

	if ok {
		if float64(timeNow()) > expiry {
			return nil, errors.New("Token is expired")
		}
	}

	return mapClaims, nil
}

func Sign(claims map[string]interface{}, naclKey string) (string, error) {

	byteClaims, err := json.Marshal(claims)

	if err != nil {
		return "", errors.New("Error in JSON marshal of claims")
	}

	var b64claims string

	b64claims, err = cryptohelper.SecretboxEncrypt(string(byteClaims), naclKey)

	if err != nil {
		return "", errors.New("Error encrypting claims")
	}

	return b64claims, nil
}

func timeNow() int64 {
	return time.Now().Unix()
}
