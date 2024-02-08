package vars

import "fmt"

// go에서 struct 는 python의 __init__랑 비슷하다고 보면 된다. 맨 초기값이라고 봐도 된다.
type Person struct {
	name    string
	age     int
	favFood []string
}

func structs() {
	// struct 만들어서 선언할 수도 있고
	favFood := []string{"buldak", "shushu"}
	cho := Person{"chocho", 25, favFood}
	shin := Person{name: "shin", age: 20, favFood: favFood}
	// slice로 선언할 때 favFood : 는 꼭 붙여주길,,
	fmt.Println(cho)
	fmt.Println(shin)
}
