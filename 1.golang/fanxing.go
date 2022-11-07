package main

import "fmt"

func test[T any](i T) T {//any 也可以写成string|int|bool
	return i
}

//type My[A any] struct {
//	Name A
//}//不允许暴露
type my[A any] struct {
	Name A
}

//类型约束
type MyType interface {
	string|int|bool
}
type myM[K MyType,V any] map[K]V

type TSlice[S comparable] []S


type MyType1 interface {
	~int|~int64  // MyInt 是由 int 来的，~int 就可以让 MyInt 也被 MyType 约束
}
type MyInt int 
func test1[T MyType1](t T) {
	fmt.Println(t)
}

func main() {
	fmt.Println(test(123))
	fmt.Println(test[string]("123"))
	fmt.Println(test[int](123)+123)
	//
	m := my[string]{
		Name: "zc",
	}
	fmt.Println(m)
	c:=make(myM[string,int])
	c["age"] = 123
	fmt.Println(c)
	//
	s := make(TSlice[string],6)
	s[5]="456 "
	//
	test1[MyInt](123)
}

