package main

import (
	"fmt"
	"os"
	"strings"
)

func decryptFile(filepath, password string) error {
	file, err := os.ReadFile("./" + filepath)
	if err != nil {
		return fmt.Errorf("Error reading file: %v\n", err)
	}

	extension := file[:4]
	nonce := file[4:16]
	key, _ := rehashPassword(password, file[16:32])

	plain, err := decrypt(key, file[32:], nonce)
	if err != nil {
		return fmt.Errorf("Error decrypting file: %v\n", err)
	}

	decryptedFile, err := os.Create("ThisOne" + "." + strings.TrimSpace(string(extension)))
	if err != nil {
		return fmt.Errorf("Error creating encrypted file: %v\n", err)
	}
	defer decryptedFile.Close()

	decryptedFile.Write(plain)

	fmt.Println("Success!")

	return nil
}
