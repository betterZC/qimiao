package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a bool = true
	fmt.Println(*(&a))
	fmt.Printf("%T", a)
	//println('66') // 单个
	str := strconv.Itoa(66)
	fmt.Println(str)
	b := 2
	switch b {
	case 2: fmt.Println("2")
	fallthrough
	default:
		fmt.Println("666")
	}
	//
	//循环体
	a1:=0
	A:
		for a1 < 10{
			a1++
			fmt.Println(a)
			if a1==5 {
				break A
				goto B
			}
		}
		
	B:
		fmt.Println("BBBBBBBB")
		//goto A
	
	//
	for{}
}
