package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	passwordLength := 16 // desired password length

	randomBytes := make([]byte, passwordLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Random Byte: ", randomBytes)
	password := base64.StdEncoding.EncodeToString(randomBytes)
	fmt.Println(password)

	fmt.Println("Generating Random Passwords using strings.random")

	rand.Seed(time.Now().UnixNano())
	password := strings.Random(16)
	fmt.Println(password)

}
