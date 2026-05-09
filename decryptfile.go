package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func decryptFile(filePath, password string) error {
	file, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	extension := file[:4]
	nonce := file[4:16]
	key, _ := rehashPassword(password, file[16:32])

	plain, err := decrypt(key, file[32:], nonce)
	if err != nil {
		return fmt.Errorf("Error decrypting file: %v", err)
	}

	splitFile := strings.Split(filePath, ".")

	decryptedFile, err := os.Create(filepath.Clean(splitFile[0]) + "-decrypted" + "." + strings.TrimSpace(filepath.Clean(string(extension))))
	if err != nil {
		return fmt.Errorf("Error creating encrypted file: %v", err)
	}
	defer decryptedFile.Close()

	_, err = decryptedFile.Write(plain)
	if err != nil {
		return fmt.Errorf("Error writing plain text to file: %v", err)
	}

	fmt.Println("Success!")

	return nil
}
