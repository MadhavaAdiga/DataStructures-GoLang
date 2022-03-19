package main

import (
	"fmt"

	"com.github/MadhavaAdiga/GolangDS/standard/list"
)

func main() {
	arrList := list.NewArrayList(10)
	arrList.Add(10)
	err := arrList.Add("qwe")

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(arrList.Size())
	}
}
