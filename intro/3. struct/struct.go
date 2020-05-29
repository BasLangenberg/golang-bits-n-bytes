package main

import "fmt"

type TestStruct struct {
	Field1 string
}

func main() {
	var ts1 TestStruct
	ts1.Field1 = "waarde1"

	var ts2 *TestStruct
	ts2 = new(TestStruct)
	ts2.Field1 = "waarde2"

	ts3 := TestStruct{
		Field1: "waarde3",
	}

	ts4 := &TestStruct{
		Field1: "waarde4",
	}

	fmt.Printf("%v | %v | %v | %v\n", ts1.Field1, ts2.Field1, ts3.Field1, ts4.Field1)

	fmt.Printf("%v | %v\n", ts1 == ts3, ts1 == *ts2)

	var ts5 TestStruct
	ts5.Field1 = "waarde1"

	var ts6 *TestStruct
	ts6 = new(TestStruct)
	ts6.Field1 = "waarde1"

	fmt.Printf("%v | %v\n", ts1 == ts5, ts1 == *ts6)
}
