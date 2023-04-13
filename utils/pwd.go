package utils

import "golang.org/x/crypto/bcrypt"

func GenHashPassword(pwd string) (hpwd string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(bytes), err
}
func CheckPassword(pwd, hpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(pwd))
	return err == nil
}
