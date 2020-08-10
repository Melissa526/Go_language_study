package someithing

import "fmt"

// export function
func SayHello() {
	fmt.Println("Hello!")
}

// non-export function
func sayBye() {
	fmt.Println("Bye~")
}
