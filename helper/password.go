package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plain string) (interface{},error) {
	password,err := bcrypt.GenerateFromPassword([]byte(plain),bcrypt.DefaultCost)
	if err != nil {
		return nil,err
	}
	return string(password),nil
}

func VerifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
