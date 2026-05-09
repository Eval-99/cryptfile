package main

import (
	"fmt"
	"mime"
	"os"
	"strings"
)

func encryptFile(filepath, password string) error {
	splitFile := strings.Split(filepath, ".")
	extension := splitFile[len(splitFile)-1]

	nonce := generateNonce(12)
	key, salt, err := hashPassword(password)

	if !strings.HasPrefix(mime.TypeByExtension("."+extension), "text/") {
		return fmt.Errorf("Invalid file type: Can only process text files\n")
	}

	file, err := os.ReadFile("./" + filepath)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	ciphertext, err := encrypt(key, file, nonce)
	if err != nil {
		return fmt.Errorf("Error encrypting file: %v", err)
	}

	newFile, err := os.Create(splitFile[0] + ".bin")
	if err != nil {
		return fmt.Errorf("Error creating encrypted file: %v", err)
	}
	defer newFile.Close()

	if len(extension) < 4 {
		extension += " "
	}

	newFile.Write([]byte(extension))
	newFile.Write(nonce)
	newFile.Write(salt)
	newFile.Write(ciphertext)

	fmt.Println("Success!")

	return nil
}
