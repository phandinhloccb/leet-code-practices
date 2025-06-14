package main

import "fmt"

func main() {
	var a int8 = -1
	var b uint8 = 255
	first := ^uint(0) >> 1

	fmt.Println("int8(-1):", a)
	fmt.Println("uint8(255):", b)
	fmt.Println("uint(0):", first)

	fmt.Printf("Binary of int8(-1): %08b\n", uint8(a))
	fmt.Printf("Binary of uint8(255): %08b\n", b)
}