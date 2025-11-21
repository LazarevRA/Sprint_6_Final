package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(s string) bool {

	// Если строка пустая - не код Морзе
	if len(s) == 0 {
		return false
	}

	// Проверяем, что строка не состоит только из пробелов
	if strings.TrimSpace(s) == "" {
		return false
	}

	// Проверяем, что строка состоит только из разрешенных символов
	return !strings.ContainsFunc(s, func(r rune) bool {
		return r != '-' && r != '.' && !unicode.IsSpace(r)
	})
}

func Convert(s string) string {
	if IsMorse(s) {
		return (morse.ToText(s))
	} else {
		return morse.ToMorse(s)
	}
}
