package testpkg

import "fmt"

var slice []int

var stringMap map[string]string

var justInt int

func testMapFunc(x map[string]string) {
	stringMap = x // want "copy-boundary: copies a map directly"
}

func testMapCopyFunc(x map[string]string) {
	stringMap = make(map[string]string, len(x))
	for key, val := range x {
		stringMap[key] = val
	}
}

func testSliceFunc(x []int) {
	slice = x // want "copy-boundary: copies a slice directly"
}

func testMultipleSliceFunc() {
	data1 := []int{1, 2, 3}
	data2 := []int{1, 2, 3}
	dataCopied1, dataCopied2 := data1, data2 // want "copy-boundary: copies a slice directly" "copy-boundary: copies a slice directly"
	fmt.Println(dataCopied1, dataCopied2)
}

func testAssignMultipleSliceFunc() {
	data1, data2 := []int{1, 2, 3}, []int{1, 3}
	dataCopied1, dataCopied2 := data1, data2 // want "copy-boundary: copies a slice directly" "copy-boundary: copies a slice directly"
	fmt.Println(dataCopied1, dataCopied2)
}

func testAssignMultipleDiffTypeFunc() {
	data1, data2 := "stringData", 5
	dataCopied1, dataCopied2 := data1, data2
	fmt.Println(dataCopied1, dataCopied2)
}

func testAssignMultipleDiffTypeWithSliceFunc() {
	data1, data2 := "stringData", []int{1, 2, 3}
	dataCopied1, dataCopied2 := data1, data2 // want "copy-boundary: copies a slice directly"
	fmt.Println(dataCopied1, dataCopied2)
}

func testSliceCopyFunc(x []int) {
	slice = make([]int, len(x))
	copy(slice, x)
}

func testEmptyAssign() {
	_ = slice
	_, newSlice := slice, []int{1, 3}
	fmt.Println(newSlice)
}

func testAssignFunc(x int) {
	slice = []int{1, 2, 2}

	newSlice := []int{1, 2, 3}
	sliceCopied := newSlice // want "copy-boundary: copies a slice directly"
	fmt.Println(sliceCopied)
}

func testNormalFunc(x int) {
	justInt = 2
	justInt = x
}

func testReturnSliceFunc() []int {
	return []int{1, 2, 3}
}

func testAssignReturnedSliceFunc(x int) {
	slice = testReturnSliceFunc()
}

func testAssignMapAndSliceFunc() {
	newMap, newSlice := []int{1, 2, 3}, map[string]string{"abc": "def"}
	anotherMap, anotherSlice := newMap, newSlice // want "copy-boundary: copies a slice directly" "copy-boundary: copies a map directly"
	fmt.Println(anotherMap, anotherSlice)
}
