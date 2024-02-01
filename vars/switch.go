package vars

import "fmt"

func canIDrinkNow(age int) bool {
	// switch의 첫번째 사용법
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 18:
		return true
	}
	return false
}

func canKoreansDrinkNow(age int) bool {
	switch korAge := age + 2; korAge {
	case 12:
		return false
	case 22:
		return true
	}
	return false
}

func showDrink() {
	fmt.Println(canIDrinkNow(20))
	fmt.Println(canKoreansDrinkNow(20))
}
