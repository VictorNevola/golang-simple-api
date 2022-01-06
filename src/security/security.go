package security

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e retorna um hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//CheckHash recebe uma string e um hash e retorna true se o hash for igual ao hash da string
func CheckHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
