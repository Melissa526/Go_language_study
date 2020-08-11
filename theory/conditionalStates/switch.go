package conditionalStates

func CanIDrinkSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}

	//if-else가 난무하는것을 막을 수 있따
	//if문처럼 variable-expression을 사용 할 수 있당
	switch {
	case age > 50:
		return false
	case age == 18:
		return true
	}

	return false
}
