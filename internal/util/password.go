package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword, düz metin şifreyi bcrypt algoritması ile hash'ler ve salt ekler.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("şifre hashlenemedi: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword, kullanıcının girdiği şifre ile veritabanındaki hash'lenmiş şifreyi karşılaştırır.
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
