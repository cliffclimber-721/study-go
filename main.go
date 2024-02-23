package main

import (
	"fmt"

	mydict "github.com/cliffclimber-721/study-go/dicts"
)

func main() {
	dictionary := mydict.Dictionary{"1": "first word"}
	defs, err := dictionary.Search("blockchain")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(defs)
	}
}
