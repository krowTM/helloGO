package main

import (
	"fmt"

	"github.com/krowTM/helloGOLib"
)

func main() {
	msg := "Hello, worlds!"
	fmt.Println(msg)
	reverseMsg, err := helloGOLib.Reverse(msg)
	if err == nil {
		fmt.Println(reverseMsg)
	}
	fmt.Println(s(4))
}

var s = func(c int) int {
	return 3 * c
}
