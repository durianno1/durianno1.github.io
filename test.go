package main

import "fmt"

func main() {
	c := []int{1, 2, 3}
	tool(&c)
	fmt.Println(c)
}

func tool(c *[]int) {
	d := c
	*d[0] = 3
	fmt.Println(d)
}
