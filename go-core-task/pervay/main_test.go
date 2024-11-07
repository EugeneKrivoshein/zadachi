package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestToStringCombined(t *testing.T) {
	expected := "4242423.140000Golangtrue(1+2i)"
	result := ToStringCombined()
	if result != expected {
		t.Errorf("ToStringCombined() = %v; expected %v", result, expected)
	}
}

func TestHashRuneWithSalt(t *testing.T) {
	// Задаем входные данные и ожидаемый результат
	input := "4242423.140000Golangtrue(1+2i)"
	salt := "go-2024"
	expectedHash := sha256.Sum256([]byte("4242423.140000Ggo-2024olangtrue(1+2i)"))
	expected := hex.EncodeToString(expectedHash[:])

	result := HashRuneWithSalt(input, salt)
	if result != expected {
		t.Errorf("HashRuneWithSalt() = %v; expected %v", result, expected)
	}
}
