package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/amolpratap-singh/file-encrypter/filecrypt"
	"github.com/amolpratap-singh/file-encrypter/logger"
	"golang.org/x/term"
)

func main() {
	logger.Log.Info().Msg("Go CLI Tool for file encryption and decryption")

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
		logger.Log.Info().Msg("Execute EncrypterCLI to encrypt or decrypt a file")
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
	if len(os.Args) < 3 {
		logger.Log.Info().Msg("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		logger.Log.Panic().Msg("File not found")
	}

	logger.Log.Info().Msg("Enter Password: ")
	password := getPassword()

	logger.Log.Info().Msg("Encrytping the file ...")
	//logger.Log.Info().Msgf("%v",password)

	filecrypt.Encrypt(file, password)
	logger.Log.Info().Msg("File encrypted succesfully ...")
}

func decryptHandler() {
	if len(os.Args) < 3 {
		logger.Log.Info().Msg("Missing the path to the file. For more information run EncrypterCLI help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		logger.Log.Panic().Msg("File not found")
	}

	logger.Log.Info().Msg("Enter Password: ")
	password := getPassword()

	logger.Log.Info().Msg("Decrytping the file ...")

	filecrypt.Decrypt(file, password)
	logger.Log.Info().Msg("File succesfully decrypted ...")
}

func validateFile(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func getPassword() []byte {
	password, _ := term.ReadPassword(0)
	logger.Log.Info().Msg("Confirm password: ")
	confirmPassword, _ := term.ReadPassword(0)

	if !validatePassword(password, confirmPassword) {
		logger.Log.Info().Msg("Passwords do not match. Please try again.")
		return getPassword()
	}

	return password
}

func validatePassword(password []byte, confirmPassword []byte) bool {
	if !bytes.Equal(password, confirmPassword) {
		return false
	}
	return true
}
