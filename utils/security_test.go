package utils

import (
	"testing"
)

func TestEncryptWithKey(t *testing.T) {
	fakePass := "fakePass"

	encryptedPass := EncryptWithKey(fakePass, []string{"super-key"})

	if encryptedPass == fakePass {
		t.Error("encrypted password is equal to fakePass")
	}
}

func TestDecryptWithKey(t *testing.T) {
	fakePass := "fakePass"

	encryptedPass := EncryptWithKey(fakePass, []string{"super-key"})

	if encryptedPass == fakePass {
		t.Error("encrypted password is equal to fakePass")
	}

	initialPass := DecryptWithKey(encryptedPass, []string{"super-key"})

	if initialPass != fakePass {
		t.Error("decrypted password is not equal to fakePass")
	}
}
