package vars

import "fmt"

func maps() {
	// map[string]인거면 입력되는 값들은 모두 문자로 선언해야하고 map[string]string은 출력값들 모두 문자여야한다.
	// key value 값이 존재하고 key 값은 대괄호에 넣고 그 옆에 value 값을 넣어준다.
	cho := map[string]string{"name": "chocho", "age": "12"}
	fmt.Println(cho)

	for key, val := range cho {
		fmt.Println(key, val)
		// val 만 출력하려고 하면
		// chocho
		// 12 이렇게 나오고
		// key 는 map[age:12 name:chocho] 이렇게 나온다
	}
}
