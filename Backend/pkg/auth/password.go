package auth

import "golang.org/x/crypto/bcrypt"

/*

Exact Process for Writing password.go:
Declare the package as auth.
Import the golang.org/x/crypto/bcrypt package.
Write HashPassword(password string) (string, error) which uses bcrypt.GenerateFromPassword to hash the password securely before saving it to the database.
Write CheckPasswordHash(password, hash string) bool which uses bcrypt.CompareHashAndPassword to verify the password during login.

*/

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
