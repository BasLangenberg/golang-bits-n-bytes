package main

import "fmt"

type TestStruct struct {
	Field1 string
}

func (ts TestStruct) PrintMe() {
	fmt.Println(ts.Field1)
}

func PrintIt(ts TestStruct) {
	fmt.Println(ts.Field1)
}

func GetTestStruct() (TestStruct, TestStruct) {
	ts1 := TestStruct{
		Field1: "waarde1",
	}
	ts2 := TestStruct{
		Field1: "waarde2",
	}

	return ts1, ts2
}

func main() {
	ts1, ts2 := GetTestStruct()
	PrintIt(ts1)
	ts2.PrintMe()
}
