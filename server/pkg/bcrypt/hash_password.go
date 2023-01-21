package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(pas string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(pas), 10)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}

func CheckPasswordHash(password, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}
