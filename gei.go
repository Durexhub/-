package main

import "fmt"

const sd = 15

func main() {

	Get("gleb", "zheg")

}
func Get(na, su string) {
	fmt.Printf("Hi %s %s \n", na, su)
}

// bool true false
// string ""
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// byte
// rune
// float32, float64 4.0
// complex64, complex128
