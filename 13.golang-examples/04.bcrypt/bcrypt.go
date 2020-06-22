package main

import (
	"bytes"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(plainText string) (hashText string) {
	preparedPlainText := preparePasswordInput(plainText)
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPlainText), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashText = string(passwordHashInBytes)
	return
}

func CompareHash(plainText string, hashText string) (err error) {
	preparedPlainText := preparePasswordInput(plainText)
	plainTextInBytes := []byte(preparedPlainText)
	hashTextInBytes := []byte(hashText)
	err = bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	return
}

func preparePasswordInput(plainText string) (preparedPasswordInput string) {
	hashedInput := sha512.Sum512_256([]byte(plainText))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPasswordInput = string(trimmedHash)
	return
}

func main() {
	password := "mypassword"

	for i := 0; i < 10; i++ {
		hashedPassword := CreateHash(password)
		err := CompareHash(password, hashedPassword)
		if err == nil {
			fmt.Printf("%s is an hash version of %s \n", hashedPassword, password)
		} else {
			fmt.Println(err.Error())
		}
	}

	userPassword := "this-is-user-password"
	userPasswordWrong := "wrong-password-input-by-user"
	userPasswordHashed := CreateHash(userPassword)

	if err := CompareHash(userPassword, userPasswordHashed); err == nil {
		fmt.Printf("You have inputed the right password!, hash[%v]\n", userPasswordHashed)
	}

	if err := CompareHash(userPasswordWrong, userPasswordHashed); err != nil {
		fmt.Printf("You have inputed the wrong password!, hash[%v]\n", userPasswordHashed)
	}
}
