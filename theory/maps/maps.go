package maps

import "fmt"

/*
Maps
- map은 javascript object와 비슷하지만 똑같지는 않다
- 선언법
variable := map[key type]value type{ initial value }

- map도 range를 이용하여 반복문에서 사용할 수 있다

*/

func testMap() {
	diue := map[string]string{
		"name" : "jisoo",
		"age" : "28" 
	}
	fmt.Println(diue)
	

	for key, value := range diue{
		fmt.Println(value)
	}
	//expected output : map[age:28 name:jisoo]
}
