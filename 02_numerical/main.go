package main

import "fmt"

func main() {
	//fmt.Printf("%d \t %b \t %x \n", 42, 42, 42)

	for i := 0; i < 200; i++ {
		fmt.Printf("%03d \t %08b \t %02x  \t %q\n", i, i, i, i)
	}
}
