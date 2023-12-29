package cmd

import (
	"os"
)

// it will related to config file like yaml where user can set the default value
// for the vm TODO SET CONTEXT LOCAL OR REMOTE here

// create a local .aspera folder in the home directory
// it will be the default location for the config file
// it will cointain the config file and the token file
// in config we save the local or remote context and credentials for local context

func initializeConfig() {

}

func getConfig() {

}

func setConfig() {
	// create the config file in the .aspera folder
	// set the default value for the context

	// create the .asper and config file
	_, err := os.Create("~/.aspera/config.yaml")
	if err != nil {
		panic(err)
	}

	// set the default value for the context
	SetContext(context)

}

func setContext(context) {

}

func init() {

}

// this function should be called only by the cli context and *CLI struct

// func Encrypt() {
// 	// encrypt the user password and save it in the config file

// 	key := []byte("passphrasewhichneedstobe32bytes!")
// 	plaintext := []byte("hello world")
// 	// encrypt the plaintext
// 	ciphertext, err := encrypt(key, plaintext)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// print the ciphertext
// 	fmt.Printf("%s", ciphertext)
// }

// func Decrypt() {

// 	key := []byte("passphrasewhichneedstobe32bytes!")
// 	// decrypt the ciphertext
// 	plaintext, err := decrypt(key, ciphertext)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// print the plaintext
// 	fmt.Printf("%s", plaintext)
// }
