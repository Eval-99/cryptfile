package main

import (
	"fmt"
	"log"
	"mime"
	"os"
	"strings"
)

func main() {
	commandArgs := os.Args
	if len(commandArgs) != 4 {
		fmt.Println("Must have '{encrypt/decrypt} {filepath} {password}' to use")
		return
	}

	action := commandArgs[1]
	filepath := commandArgs[2]
	password := commandArgs[3]

	switch action {
	case "encrypt":
		splitFile := strings.Split(filepath, ".")
		extension := splitFile[len(splitFile)-1]

		nonce := generateNonce(12)
		key, salt, err := hashPassword(password)

		if !strings.HasPrefix(mime.TypeByExtension("."+extension), "text/") {
			fmt.Println("Invalid file type: Can only process text files")
			return
		}

		file, err := os.ReadFile("./" + filepath)
		if err != nil {
			log.Fatal("Error reading file")
		}

		ciphertext, err := encrypt(key, file, nonce)
		if err != nil {
			log.Fatal("Error encrypting file")
		}

		newFile, err := os.Create(splitFile[0] + ".bin")
		if err != nil {
			log.Fatal("Error creating encrypted file")
		}
		defer newFile.Close()

		if len(extension) < 4 {
			extension += " "
		}

		newFile.Write([]byte(extension))
		newFile.Write(nonce)
		newFile.Write(salt)
		newFile.Write(ciphertext)
	case "decrypt":
		file, err := os.ReadFile("./" + filepath)
		if err != nil {
			log.Fatal("Error reading file")
		}

		extension := file[:4]
		nonce := file[4:16]
		key, _ := rehashPassword(password, file[16:32])

		plain, err := decrypt(key, file[32:], nonce)
		if err != nil {
			fmt.Println(err)
		}

		decryptedFile, err := os.Create("ThisOne" + "." + strings.TrimSpace(string(extension)))
		if err != nil {
			log.Fatal("Error creating encrypted file")
		}
		defer decryptedFile.Close()

		decryptedFile.Write(plain)
	}
}
