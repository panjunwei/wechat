package util

import "github.com/panjunwei/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
