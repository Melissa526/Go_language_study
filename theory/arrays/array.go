package arrays

import "fmt"

func GetNames() {
	// Array선언하는법
	// variable := [length]type{initial value}
	//ex)
	names := [5]string{"nico", "ana", "jamie", "elly", "bin"}
	fmt.Println(names)

	//unlimited array 선언 = slice
	namess := []string{"nico", "ana", "jamie"}
	newnames := append(namess, "flynn") //namess에 직접 추가되는게 아님, 새로운 slice를 return
	fmt.Println(namess)
	fmt.Println(newnames)
}
