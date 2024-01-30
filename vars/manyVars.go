package vars

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}

func ShowRepeat() {
	repeatMe("chocho", "shin", "kim", "park")
}
