package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func encryptFile(filePath, password string) error {
	splitFile := strings.Split(filePath, ".")
	extension := splitFile[len(splitFile)-1]

	if len(extension) > 4 {
		return fmt.Errorf("File extension length must be at most 4 characters long")
	}

	nonce := generateNonce(12)
	key, salt, err := hashPassword(password)

	if !strings.HasPrefix(mime.TypeByExtension("."+extension), "text/") {
		return fmt.Errorf("Invalid file type: Can only process text files")
	}

	file, err := os.ReadFile(filepath.Clean(splitFile[0]))
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	ciphertext, err := encrypt(key, file, nonce)
	if err != nil {
		return fmt.Errorf("Error encrypting file: %v", err)
	}

	newFile, err := os.Create(filepath.Clean(splitFile[0]) + ".bin")
	if err != nil {
		return fmt.Errorf("Error creating encrypted file: %v", err)
	}
	defer newFile.Close()

	if len(extension) < 4 {
		extension += " "
	}

	_, err = newFile.Write([]byte(extension))
	if err != nil {
		return fmt.Errorf("Error writing extension: %v", err)
	}
	_, err = newFile.Write(nonce)
	if err != nil {
		return fmt.Errorf("Error writing nonce: %v", err)
	}
	_, err = newFile.Write(salt)
	if err != nil {
		return fmt.Errorf("Error writing salt: %v", err)
	}
	_, err = newFile.Write(ciphertext)
	if err != nil {
		return fmt.Errorf("Error writing ciphertext: %v", err)
	}

	fmt.Println("Success!")

	return nil
}
