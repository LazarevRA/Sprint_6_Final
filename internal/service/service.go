package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(s string) bool {
	//Если строка пустая - не код Морзе
	if len(s) == 0 {
		return false
	}

	morseStatus := false

	for _, ch := range s {
		//Проверяем, что строка состоит только из разрешенных символов
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
		//Проверяем, что строка не состоит только из пробелов
		if ch == '.' || ch == '-' {
			morseStatus = true
		}
	}
	return morseStatus
}

func Convert(s string) string {
	if IsMorse(s) {
		return (morse.ToText(s))
	} else {
		return morse.ToMorse(s)
	}
}
