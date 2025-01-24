package app

import (
	"math/rand"
	"time"
	"gopass/config"
	"gopass/helper"
	
)

const (
	Lower = "abcdefghijklmnopqrstuvwxyz"
	Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers = "0123456789"
	Symbols = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func GeneratePassword(cfg config.Config) (string, error) {
	if cfg.Length <= 0 {
		return "", nil
	}

	charPool := ""
	if cfg.IncludeLower {
		charPool += Lower
	}
	if cfg.IncludeUpper {
		charPool += Upper
	}
	if cfg.IncludeNumbers {
		charPool += Numbers
	}
	if cfg.IncludeSymbols {
		charPool += Symbols
	}
	if charPool == "" {
		return "", nil
	}

	rand.Seed(time.Now().UnixNano())

	password := make([]byte, cfg.Length)
	for i := range password {
		password[i] = charPool[rand.Intn(len(charPool))]
	}

	return helper.ShuffleString(string(password)), nil
}
