package main

import (
	"crypto/sha256"
	"fmt"
)

type HashTable struct {
	size  int
	table [][]interface{}
}

func (h *HashTable) hash(key string) string {
	sum := sha256.Sum256([]byte("hello world"))
	return fmt.Sprintf("%x", sum)
}

func main() {
}
