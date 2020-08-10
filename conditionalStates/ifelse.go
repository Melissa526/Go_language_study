package conditionalStates

func CanIDrink(age int) bool {
	//Go Extension에서는 else문 안써도 if문이틀리면 자동으로 else로 동작
	if age < 18 {
		return false
	}
	return true
	/* same code
	if age < 18 {
		return false
	} else {
		return true
	} */
}

func CanIDrinkInKorea(age int) bool {
	//variable expression
	//if문 선언과 동시에 변수 선언가능
	//if문에서 세미콜론(;)을 기준으로 변수를 선언하고 조건문을 사용할 수 있다

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
