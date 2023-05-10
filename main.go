package main

import (
	"fmt"
)

func main() {
	fmt.Println(1<<0, 1<<1, 1<<2, 1<<3, 1<<4, 1<<32, 1<<62)

	fmt.Println(1<<1 | 1<<3)

	fmt.Println("🐟 Fishing...")
}
