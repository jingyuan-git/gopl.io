package main

import (
	"fmt"
	"github.com/k-sone/critbitgo"
)

func main() {
	// Create Trie
	trie := critbitgo.NewTrie()

	// Insert
	trie.Insert([]byte("aa"), "value1")
	trie.Insert([]byte("bb"), "value2")
	trie.Insert([]byte("ab"), "value3")

	// Get
	v, ok := trie.Get([]byte("aa"))
	fmt.Println(v, ok) // -> value1 true

	// Iterate containing keys
	trie.Allprefixed([]byte{}, func(key []byte, value interface{}) bool {
		fmt.Println(key, value) // -> [97 97] value1
		//    [97 98] value3
		//    [98 98] value2
		return true
	})

	// Delete
	v, ok = trie.Delete([]byte("aa"))
	fmt.Println(v, ok) // -> value1 true
	v, ok = trie.Delete([]byte("aa"))
	fmt.Println(v, ok) // -> <nil> false
}
