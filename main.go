package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello", ReturnWorld())
}

func ReturnWorld() string {
	return "World!!!"
}
