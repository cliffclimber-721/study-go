package vars

import (
	"fmt"
)

func superAdd(nums ...int) int {
	// C언어에서 쓰듯 for 문을 제시할 수 있다.
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	return 1
}

func superShowAdd(numbers ...int) int {
	// range를 붙이면 인덱스와 숫자 둘다 부여하게 돼서 superShowAdd가 출력하게 되면 값이 0부터 나온다.
	// index 안 쓰고 싶으면 _ 써서 없앨 수 있다.
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func Result() {
	//superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	res := superShowAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)
}
