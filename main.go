package main

import (
	"fmt"
	"os"
)

func main() {
	commandArgs := os.Args
	if commandArgs[1] == "help" {
		fmt.Println(`
Action:   Use the encrypt or decrypt action
Filepath: Location of the file
Password: Key to lock and unlock file

Example:
	  encrypt somefile.txt password
	  decrypt somefile.bin password
	`)

		return
	}
	if len(commandArgs) != 4 {
		fmt.Println("Must have '{encrypt/decrypt} {filepath} {password}' to use")
		return
	}

	action := commandArgs[1]
	filepath := commandArgs[2]
	password := commandArgs[3]

	switch action {
	case "encrypt":
		err := encryptFile(filepath, password)
		if err != nil {
			fmt.Println(err)
		}

	case "decrypt":
		err := decryptFile(filepath, password)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Not a valid action: Must encrypt or decrypt")
	}
}
