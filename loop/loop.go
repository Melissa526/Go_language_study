package loop

import "fmt"

//Go에서는 for만 사용가능(foreach, for in, for of X)
func SuperAdd(numbers ...int) int {

	/* 1. for 일반적인 형태*/
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	/* 2.forEach같은형태 */
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	/* Example */
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
