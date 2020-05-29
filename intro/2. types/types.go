package main

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

var (
	// string types
	s string = "Hello world"

	// numeric types
	u8  uint8  = math.MaxUint8
	u16 uint16 = math.MaxUint16
	u32 uint32 = math.MaxUint32
	u64 uint64 = math.MaxUint64

	i8  int8  = math.MaxInt8
	i16 int16 = math.MaxInt16
	i32 int32 = math.MaxInt32
	i64 int64 = math.MaxInt64

	f32 float32 = math.MaxFloat32
	f64 float64 = math.MaxFloat64

	c64  complex64  = complex(math.MaxFloat32, math.MaxFloat32)
	c128 complex128 = complex(math.MaxFloat64, math.MaxFloat64)

	// bool types
	b bool = true

	// error types
	err error = errors.New("Foutje bedankt")
)

func main() {
	fmt.Printf("Type: %T Length: %d Value: %v\n", s, unsafe.Sizeof(s), s)

	fmt.Printf("Type: %T Length: %d Value: %v\n", u8, unsafe.Sizeof(u8), u8)
	fmt.Printf("Type: %T Length: %d Value: %v\n", u16, unsafe.Sizeof(u16), u16)
	fmt.Printf("Type: %T Length: %d Value: %v\n", u32, unsafe.Sizeof(u32), u32)
	fmt.Printf("Type: %T Length: %d Value: %v\n", u64, unsafe.Sizeof(u64), u64)

	fmt.Printf("Type: %T Length: %d Value: %v\n", i8, unsafe.Sizeof(i8), i8)
	fmt.Printf("Type: %T Length: %d Value: %v\n", i16, unsafe.Sizeof(i16), i16)
	fmt.Printf("Type: %T Length: %d Value: %v\n", i32, unsafe.Sizeof(i32), i32)
	fmt.Printf("Type: %T Length: %d Value: %v\n", i64, unsafe.Sizeof(i64), i64)

	fmt.Printf("Type: %T Length: %d Value: %v\n", f32, unsafe.Sizeof(f32), f32)
	fmt.Printf("Type: %T Length: %d Value: %v\n", f64, unsafe.Sizeof(f64), f64)

	fmt.Printf("Type: %T Length: %d Value: %v\n", c64, unsafe.Sizeof(c64), c64)
	fmt.Printf("Type: %T Length: %d Value: %v\n", c128, unsafe.Sizeof(c128), c128)

	fmt.Printf("Type: %T Length: %d Value: %v\n", b, unsafe.Sizeof(b), b)

	fmt.Printf("Type: %T Length: %d Value: %v\n", err, unsafe.Sizeof(err), err)

	var s1 string = "Hello world"
	var s2 string
	s2 = "Hello world 2"

	var (
		s3 string = "Hello world the 3rd"
		s4 string
	)

	s4 = "Hello world 4"

	const s5 = "Hello 5 worlds"

	fmt.Printf("%v | %v | %v | %v | %v\n", s1, s2, s3, s4, s5)

	i1 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", i1)

	var i2 [6]int = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", i2)
}
