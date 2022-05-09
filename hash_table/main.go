package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type HashTable struct {
	size int64
	// To DO : What is the differences between [][]interface and [][]map[interface{}]intreface{}
	// table [][]interface{}
	table [][]map[interface{}]interface{}
}

func newHashTable(size int64) *HashTable {
	return &HashTable{size: size, table: make([][]map[interface{}]interface{}, size)}
}

// To Do: Consider the type of key. string or interface{}
func (h *HashTable) hash(key string) int64 {
	sum := sha256.Sum256([]byte(key))
	parsed_sum_base_16, _ := strconv.ParseInt(fmt.Sprintf("%x", sum), 16, 64)
	return parsed_sum_base_16 % h.size
}

func (h *HashTable) add(key string, value interface{}) {
	index := h.hash(key)

	for _, data := range h.table[index] {
		if data[0] == key {
			data[1] = value
			return
		}
	}
	element := make(map[interface{}]interface{})
	element[key] = value
	h.table[index] = append(h.table[index], element)
}

func main() {
	// hashメソッドの返り値がすべて7になってしまう

	// h := Human{Name: "Bob", Age: 20}
	// s := "My name is Kazuki"

}
