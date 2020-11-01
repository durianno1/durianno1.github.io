package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	nums := make([][]string, 0)
	inputReader := bufio.NewScanner(os.Stdin)
	i := 0
	t := 0
	for inputReader.Scan() {
		cur := strings.Split(inputReader.Text(), " ")
		nums = append(nums, cur)
		if i == 0 {
			i = len(cur)
		}
		if t == i-1 {
			break
		}
		t++
	}
	fmt.Println(nums)
}
