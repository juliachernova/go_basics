package main

// Можно импортировать и так
// import "fmt"
// import "math"
// Но лучше использовать "factored" стиль импортирования пакетов

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
