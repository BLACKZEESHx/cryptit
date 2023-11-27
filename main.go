package main

import (
	"endecry/decrypt"
	"endecry/encrypt"
	"fmt"
)

func main() {
	var enstr string
	fmt.Print("Enter Text for Encrypt: ")
	fmt.Scanf("%s", &enstr)

	encrypted := encrypt.Nimbus(enstr)
	fmt.Println("Encrypted: ", encrypted)
	
	decrypted := decrypt.Nimbus(encrypted)
	fmt.Println("decrypted: ", decrypted)
}
