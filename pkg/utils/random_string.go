package utils

import (
	"math/rand"
	"strings"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func String(count int) string {
	var sb strings.Builder

	for len(sb.String()) != count {
		randomNumber := rand.Intn(len(letters))
		sb.WriteByte(letters[randomNumber])
	}

	return sb.String()
}
