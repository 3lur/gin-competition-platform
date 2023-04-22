package hash

import "golang.org/x/crypto/bcrypt"

func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsHashed(str string) bool {
	return len(str) == 60
}
