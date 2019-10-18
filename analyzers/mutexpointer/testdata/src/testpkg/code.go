package testpkg

import (
	"sync"
)

var lock *sync.Mutex // want "mutex-pointer: var lock uses pointer to sync.Mutex"
var x, y *sync.Mutex // want "mutex-pointer: var x uses pointer to sync.Mutex" "mutex-pointer: var y uses pointer to sync.Mutex"

var (
	invalidMutextInBlock *sync.Mutex // want "mutex-pointer: var invalidMutextInBlock uses pointer to sync.Mutex"
	validMutex           sync.Mutex
)

type rwm *sync.RWMutex // want "mutex-pointer: type rwm uses pointer to sync.RWMutex"

func testFunc() {
	var anotherMutexPointer *sync.Mutex // want "mutex-pointer: var anotherMutexPointer uses pointer to sync.Mutex"
	_ = anotherMutexPointer
	alsoInvalidMutex, valid, notValidToo := new(sync.Mutex), sync.Mutex{}, new(sync.RWMutex) // want "mutex-pointer: var alsoInvalidMutex uses pointer to sync.Mutex" "mutex-pointer: var notValidToo uses pointer to sync.RWMutex"
	_ = alsoInvalidMutex
	_ = valid
	_ = notValidToo
}
