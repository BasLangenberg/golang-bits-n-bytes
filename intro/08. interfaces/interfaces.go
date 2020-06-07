package main

import "fmt"

type TestInterface interface {
	MethodOne()
}

type TestStructOne struct {
}

type TestStructTwo struct {
}

func (ts TestStructOne) MethodOne() {}

func GetInstance() TestInterface {
	return TestStructOne{}
}

func main() {
	ti1 := GetInstance()
	_, ok := ti1.(TestStructOne)
	if ok {
		fmt.Println("ti1 instanceof TestStructOne")
	} else {
		fmt.Println("ti1 no instanceof TestStructOne")
	}

	// _, ok := ti1.(TestStructTwo)
	// if ok {
	// 	fmt.Println("ti1 instanceof TestStructTwo")
	// } else {
	// 	fmt.Println("ti1 no instanceof TestStructTwo")
	// }
}
