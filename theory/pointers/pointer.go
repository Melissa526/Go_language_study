package pointers

import "fmt"

//low level programing
//메모리 주소 참조
func CopyWithMemory() {
	a := 2
	b := a  //copy with value
	c := &a //메모리 주소 복사
	d := *c //see through value of memory
	a = 5
	fmt.Println(&a, &b)     // pointer(&연산자)를 쓰면 메모리 주소 볼수 있다
	fmt.Println(a, b, c, d) //expected output : 5 2 -x어쩌구 2
}

/* Summary

& : Address of memory
* : See through value of the memory

- *variable 을 통해서 값을 업데이트 할 수 있다.
   메모리를 바로 참조하기때문에 메모리성능에 좋다

*/
