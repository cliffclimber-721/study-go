package vars

import "fmt"

func nameSlice() {
	names := [5]string{"chocho", "shin", "kim"}
	names[3] = "kang"
	names[4] = "lim"
	fmt.Println(names)

	// 굳이 [] 안에 숫자를 쓸 필요는 없다
	firstname := []string{"chocho", "shin", "kim"}
	fmt.Println(firstname)
	// 이걸 slice 라고 한다.
	firstname = append(firstname, "lala")
	fmt.Println(firstname)
}
