package main

import (
	"fmt"
	"math/rand"
	// По соглашению, имя пакета такое же как и последнего элемента в import.
	// "math/rand" пакет использует rand.
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
