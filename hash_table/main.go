package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type HashTable struct {
	size  int
	table [][]interface{}
}

func (h *HashTable) hash(key string, size int64) int64 {
	sum := sha256.Sum256([]byte(key))
	parsed_sum_base_16, _ := strconv.ParseInt(fmt.Sprintf("%x", sum), 16, 64)
	return parsed_sum_base_16 % size
}

func main() {
	h := &HashTable{size: 10}
	fmt.Println(h.hash("test", 10))
}
