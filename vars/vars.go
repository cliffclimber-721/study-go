package vars

import "fmt"

// 모든 변수에 타입을 지정해줘야 한다.
func ShowVar() {
	// := 는 변수 선언할 때도 사용된다.
	name := "chocho" // var name string = "chocho" 랑 같은 뜻이다. := 는 축약된 형태를 말한다.
	name = "shin"
	fmt.Println(name)
}

// return 값에도 타입 지정을 무조건 해줘야한다.
// 소괄호 안에 고정적으로 쓰는 변수에도 무조건 타입을 지정해줘야한다.
func ShowNums(a int, b int) int {
	return a * b
}
