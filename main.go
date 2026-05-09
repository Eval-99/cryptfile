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
		fmt.Println("Must have \"{encrypt/decrypt} {filePath} {password}\" to use. See \"help\" command.")
		return
	}

	action := commandArgs[1]
	filePath := commandArgs[2]
	password := commandArgs[3]

	switch action {
	case "encrypt":
		err := encryptFile(filePath, password)
		if err != nil {
			fmt.Println(err)
		}

	case "decrypt":
		err := decryptFile(filePath, password)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Not a valid action: Must be encrypt or decrypt")
	}
}
