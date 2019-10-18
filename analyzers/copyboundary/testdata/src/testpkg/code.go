package testpkg

var slice []int

var stringMap map[string]string

var justInt int

func testMapFunc(x map[string]string) {
	stringMap = x // want "copy-slice-map: copies a map directly"
}

func testSliceFunc(x []int) {
	slice = x // want "copy-slice-map: copies a slice directly"
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
