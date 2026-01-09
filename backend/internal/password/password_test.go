package password

import (
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "my_secret_password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify format
	if !strings.HasPrefix(hash, "$argon2id$v=19$m=65536,t=1,p=4$") {
		t.Errorf("Hash format incorrect: %s", hash)
	}
}

func TestCheckPassword(t *testing.T) {
	password := "my_secure_password_123"

	// Create hash
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// 1. Correct password should match
	match, err := CheckPassword(password, hash)
	if err != nil {
		t.Errorf("CheckPassword returned error for valid password: %v", err)
	}
	if !match {
		t.Error("CheckPassword returned false for valid password")
	}

	// 2. Wrong password should fail
	match, err = CheckPassword("wrong_password", hash)
	if err != nil {
		t.Errorf("CheckPassword returned error for invalid password: %v", err)
	}
	if match {
		t.Error("CheckPassword returned true for invalid password")
	}

	// 3. Modified hash should fail
	// Tamper with the hash part
	badHash := hash[:len(hash)-5] + "AAAAA"
	match, err = CheckPassword(password, badHash)
	if err == nil && match {
		t.Error("CheckPassword should fail for tampered hash")
	}
}
