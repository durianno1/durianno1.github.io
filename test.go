package main

import (
	"fmt"
	"reflect"
)

func main() {
	type a int
	var s a = 5
	v := reflect.ValueOf(&s)
	v = v.Elem()
	fmt.Println(v.Interface())
}
