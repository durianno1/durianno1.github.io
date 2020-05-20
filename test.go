package main

import "fmt"

func main() {
	res := new(map[int]int)
	(*res)[0] = 1
	fmt.Println(*res)

}
