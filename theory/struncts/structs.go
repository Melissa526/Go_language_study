package struncts

import "fmt"

/* Structs

- class + Object를 합쳐놓은 것 특성의 (사실 Go는 class는 없당)
- object
- struct는 func(혹은 메소드)를 포함할 수 있따. => constructor가 없음
- 선언 : 함수위에 struct를 먼저 선언한다
type 'name' struct { variable type ... }

*/
type person struct{
	name	string
	age		int
	favFood	[]string 
}

func setPersion(){
	favFood := []string{"kimchi", "burger", "steak"}
	//zsoo := person{"jisoo", 28, favFood} 				//이렇게 축약으로 쓸수 도 있다
	zsoo := person{name:"jisoo", age:28, favFood:favFood}
	fmt.Println(zsoo.name)
}