package main

/**
 * Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
 * To create an empty map, use the builtin make: make(map[key-type]val-type).
 * Set key/value pairs using typical name[key] = val syntax.
 * Printing a map with e.g. fmt.Println will show all of its key/value pairs.
 * Get a value for a key with name[key].
 * If the key doesn’t exist, the zero value of the value type is returned.
 * The builtin len returns the number of key/value pairs when called on a map.
 * The builtin delete removes key/value pairs from a map.
 * To remove all key/value pairs from a map, use the clear builtin.
 * The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
 * You can also declare and initialize a new map in the same line with this syntax.
 * The maps package contains a number of useful utility functions for maps.
 *
Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
*/

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
