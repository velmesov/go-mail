package error

import "log"

// CheckNil проверка на ошибку
func CheckNil(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// CheckFlag Проверка аргументов
func CheckFlag(flagVal string, msgErr string) {
	if flagVal == "" {
		log.Fatal(msgErr)
	}
}
