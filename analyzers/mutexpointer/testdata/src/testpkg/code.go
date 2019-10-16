package pkg

import (
	"sync"
)

var lock *sync.Mutex // want "mutex-pointer: var lock uses pointer to sync.Mutex"
var x, y *sync.Mutex  // want "mutex-pointer: var x uses pointer to sync.Mutex"

var (
	invalidMutextInBlock *sync.Mutex // want "mutex-pointer: var invalidMutxtInBlock uses pointer to sync.Mutex"
	validMutex            sync.Mutex
)

type rwm *sync.RWMutex // want "mutex-pointer: var rwm points to pointer of sync.Mutex"

func testFunc() {
	var anotherMutexPointer *sync.Mutex // want "mutex-pointer: var anotherMutexPointer uses pointer to sync.Mutex"
	_ = anotherMutexPointer
	alsoInvalidMutex, valid, notValidToo := new(sync.Mutex), sync.Mutex{}, new(sync.RWMutex)
	_ = alsoInvalidMutex
	_ = valid
	_ = notValidToo
}
