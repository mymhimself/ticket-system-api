package account

import (
	"bytes"
	random "crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func createHashedPassword(password, salt string) []byte {
	hash := sha512.New()
	hash.Write([]byte(password + salt))
	return hash.Sum(nil)
}

func comparePassword(password string, storedHashedPassword []byte, storedSalt string) bool {
	hashedPassword := createHashedPassword(password, storedSalt)
	return bytes.Compare(hashedPassword, storedHashedPassword) == 0
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := random.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
