package vars

import (
	"fmt"
	"strings"
)

func lenUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 모든 변수에 타입을 지정해줘야 한다.
func ShowLenUpper() {
	totalLen, upName := lenUpper("chocho")
	fmt.Println(totalLen, upName)
	// 변수 하나를 무시하고 값을 출력할 수 있다. 이전 함수에서 return 값을 2개 줬기 때문에 무조건 2개 나오는 게 정상인데 하나만 출력하고 싶다면 _ 를 쓴다.
	totalLen2, _ := lenUpper("shin")
	fmt.Println(totalLen2)
}
