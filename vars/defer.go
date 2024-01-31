package vars

import (
	"fmt"
	"strings"
)

func lenIsUpper(name string) (length int, uppercase string) {
	// defer는 함수가 끝났을 때 출력해주는 문법이다. return 이랑 같이 써야한다!
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	// return 값을 꼭 적어서 낼 필요는 없다
	// naked return 이라고 한다
	return
}
