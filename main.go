package main

import (
	"fmt"

	"github.com/diue/learngo/conditionalStates"
)

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
func main() {
	//(1.6)
	fmt.Println(conditionalStates.CanIDrink(28))
	//(1.5)total := loop.SuperAdd(1, 2, 3, 4, 5, 6)
}
