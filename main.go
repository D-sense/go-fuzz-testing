package main

import (
	"go-fuzzing/caesar"
	"log"
)

func main() {
	// =========================================================================
	// Encoding Text
	text := "Let's carve him as a dish fit for the gods."
	key := 3
	encData := caesar.NewEncoding{
		Text: text,
		Key:  key,
	}

	encodedStr, err := caesar.Encrypter(encData)
	if err != nil {
		log.Fatalf("error encoding a text: %v", err)
	}

	log.Println("encoded text: ", encodedStr)
	log.Println("================================================")

	// =========================================================================
	// Decoding Text
	decodedData := caesar.NewDecoding{
		Text: encodedStr,
		Key:  key,
	}

	decodedStr, err := caesar.Decrypter(decodedData)
	if err != nil {
		log.Fatalf("error encoding a text: %v", err)
	}

	log.Println("decoded text: ", decodedStr)
}
