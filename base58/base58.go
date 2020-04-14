package main

import (
	"base58"
)

func main() {
	base58.Encode([]byte{100, 1})
}
