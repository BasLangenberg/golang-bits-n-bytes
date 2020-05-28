package main

import "fmt"

type TestStruct struct {
	Field1 string
}

func Method(ts TestStruct) {
	ts.Field1 = "Method"
}

func PointyMethod(ts *TestStruct) {
	ts.Field1 = "PointyMethod"
}

func main() {
	ts1 := TestStruct{
		Field1: "test1",
	}
	Method(ts1)
	fmt.Println(ts1.Field1)

	ts2 := TestStruct{
		Field1: "test2",
	}
	PointyMethod(&ts2)
	fmt.Println(ts2.Field1)
}
