// main.go 는 컴파일을 하기 위한 파일이다.

package main

import (
	"fmt"
)

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("chocho", "shin", "kim", "park")
}
