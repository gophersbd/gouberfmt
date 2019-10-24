package testpkg

import "fmt"

func testZeroSizedChanFunc() {
	c := make(chan int)
	fmt.Println(c)
}

func testOneSizedChanFunc() {
	c := make(chan int, 1)
	fmt.Println(c)
}

func testGreaterThanOneSizedChanFunc() {
	c := make(chan int, 64) // want "chan-size: channel size should be one or none"
	fmt.Println(c)
}

func testAssignFunc() {
	x, c := "justString", make(chan int, 64) // want "chan-size: channel size should be one or none"
	fmt.Println(c, x)
}
