package pkg

var ifacePointer *interface{} // want "interface-pointer: var ifacePointer uses pointer to interface"
var x, y *interface{}         // want "interface-pointer: var x uses pointer to interface" "interface-pointer: var y uses pointer to interface"
var (
	ifacePointerInBlock *interface{} // want "interface-pointer: var ifacePointerInBlock uses pointer to interface"
	validIface          interface{}
	validVar            int
)

type (
	constantIfacePointer *interface{}
)

func testFunc() {
	var anotherIfacePointer *interface{} // want "interface-pointer: var anotherIfacePointer uses pointer to interface"
	_ = anotherIfacePointer
}
