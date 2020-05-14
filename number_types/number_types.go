package main

import "fmt"

func main() {
	var i int8 = 20
	var f float32 = 5.6
	fmt.Println(float32(i) + f)
	fmt.Println(i + int8(f))
	fmt.Println(i + int8(f+1.9))

	var j int32 = 40
	fmt.Println(int32(i) + j)
	fmt.Println(i + int8(j))
}
