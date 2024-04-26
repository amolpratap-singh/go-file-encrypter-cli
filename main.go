package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go CLI Tool for file encryption and Decryption")

	if len(os.Args) < 2 {
		helper()
		os.Exit(0)
	}

	action := os.Args[1]

	switch action {
	case "help":
		helper()
	case "encrypt":
		encryptHandler()
	case "decrypt":
		decryptHandler()
	default:
		fmt.Println("Run CryptoGo encrypt to encrypt a file, and CryptoGo decrypt to decrypt a file.")
		os.Exit(1)
	}

}

func helper() {
	fmt.Println("CryptoGo")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tCryptoGo encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encryptHandler() {
	fmt.Println("Encryption of File will be done")
}

func decryptHandler() {
	fmt.Println("Decryption of File will be done")
}