package vars

import "fmt"

// 파이썬과 다르게 if문 쓸 때 소괄호나 : 가 필요없다. 그냥 조건만 적어주면 된다.
func canIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

func korCanIDrink(age int) bool {
	// 변수 선언할 때 ;를 넣어서 선언 가능하다. 아래 예시를 보면 알 수 있다.
	if korAge := age + 2; korAge < 18 {
		return false
	}
	return true
}
func main() {
	fmt.Println(canIDrink(20))
	fmt.Println(korCanIDrink(15))
}
