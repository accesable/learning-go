package auth

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func ComparePassword(hashed string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), [](byte)(password))
	if err != nil {
		return err
	}
	return err
}
