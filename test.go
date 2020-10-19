package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func main(){
	a := make([]string,2)
	scanner := bufio.NewScanner(os.Stdin)
	for i:=0;i<2;i++ {
		scanner.Scan()
		a[i] = scanner.Text()
	}
	fmt.Println(a)
}
*/
func main() {
	var rows int
	fmt.Scanln(&rows)
	//nums := make([]string,rows)
	nums := make([][]int, rows)
	inputReader := bufio.NewScanner(os.Stdin)
	//for i:=0;i<rows;i++{
	//nums[i],_ = inputReader.ReadString('\n')
	//nums[i] = strings.Replace(nums[i]," ","",-1) 删除空格
	//fmt.Println(strings.Split(inputReader.Text()," "))
	//}
	index := 0
	for inputReader.Scan() {
		for _, j := range strings.Split(inputReader.Text(), " ") {
			cur, _ := strconv.Atoi(j)
			nums[index] = append(nums[index], cur)
		}
		if index == rows-1 {
			break
		}
		index++
	}
	fmt.Println(nums)
}
