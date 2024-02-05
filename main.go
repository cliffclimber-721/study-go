package main

import "fmt"

func main() {
	s, m, l := fmt.Scan(&s, &m, &l)
	if res := 1*s + 2*m + 3*l; res > 9 {
		fmt.Println("happy")
	}
	fmt.Println("sad")
}
