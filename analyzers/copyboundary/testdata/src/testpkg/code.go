package testpkg

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

func testSliceCopyFunc(x []int) {
	slice = make([]int, len(x))
	copy(slice, x)
}

func testAssignFunc(x int) {
	slice = []int{1, 2, 2}
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
