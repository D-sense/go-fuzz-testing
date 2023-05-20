package caesar_test

import (
	"go-fuzzing/caesar"
	"testing"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

// Fuzz_Encrypter Fuzz-testing detects the bug in our Encrypter() implementation (of which simple testing could not detect)
func Fuzz_Encrypter(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string, key int) {
		data := caesar.NewEncoding{
			Text: text,
			Key:  key,
		}

		_, err := caesar.Encrypter(data)
		if err != nil {
			t.Logf("\t%s\tShould return error for a key zero value", failed)
		}

	})
}

func Test_Encrypter(t *testing.T) {
	t.Log("Encrypt text")
	{
		t.Logf("\tWhen handling a key of zero value")
		{
			data := caesar.NewEncoding{
				Text: "Let's carve him as a dish fit for the gods.",
				Key:  0,
			}

			_, err := caesar.Encrypter(data)
			if err == nil {
				t.Logf("\t%s\tShould return error for a key zero value", failed)
			}

			t.Logf("\t%s\tShould not be able to encrpt a text", success)
		}

		t.Logf("\tWhen handling a key of value 1 or greater")
		{
			data := caesar.NewEncoding{
				Text: "Let's carve him as a dish fit for the gods.",
				Key:  14,
			}

			encodedStr, err := caesar.Encrypter(data)
			if err != nil {
				t.Logf("\t%s\tShould not return error for a key greater than zero | Error:= %v", failed, err)
			}

			expectedStr := "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg."
			if encodedStr != expectedStr {
				t.Logf("\t\tExp: %s", expectedStr)
				t.Logf("\t\tgot: %s", encodedStr)
				t.Fatalf("\t%s\tShould return a valid encrypted string", failed)
			}

			t.Logf("\t%s\tShould encrypt text", success)
		}
	}
}

func Test_Decrypter(t *testing.T) {
	t.Log("Decrypt text")
	{
		t.Logf("\tWhen handling a key of zero value")
		{
			data := caesar.NewDecoding{
				Text: "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.",
				Key:  0,
			}

			encodedStr, err := caesar.Decrypter(data)
			if err == nil {
				t.Logf("\t%s\tShould return error for a key zero value", failed)
			}

			if encodedStr != "" {
				t.Logf("\t\tExp: %s", "")
				t.Logf("\t\tgot: %s", encodedStr)
				t.Fatalf("\t%s\tShould return empty string", failed)
			}

			t.Logf("\t%s\tShould decrypt text", success)
		}

		t.Logf("\tWhen handling a key of value 1 or greater")
		{
			data := caesar.NewDecoding{
				Text: "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.",
				Key:  14,
			}

			decodedStr, err := caesar.Decrypter(data)
			if err != nil {
				t.Logf("\t%s\tShould not return error for a key greater than zero", failed)
			}

			expectedStr := "Let's carve him as a dish fit for the gods."
			if decodedStr != expectedStr {
				t.Logf("\t\tExp: %s", expectedStr)
				t.Logf("\t\tgot: %s", decodedStr)
				t.Fatalf("\t%s\tShould return a valid decrypted text", failed)
			}

			t.Logf("\t%s\tShould decrypt text", success)
		}
	}
}
