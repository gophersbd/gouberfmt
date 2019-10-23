# gouberfmt

Experimental golang static analysis tool that supports - https://github.com/uber-go/guide/blob/master/style.md

## Supported Analyzers

* Pointer to Interface

    ```golang
    // interface-pointer: var ifacePointer uses pointer to interface
    var ifacePointer *interface{}
    ```
* Pointer to sync.Mutex

    ```golang
    // "mutex-pointer: var lock uses pointer to sync.Mutex"
    var lock *sync.Mutex
    ```
  
* Copy Boundary

    ```golang
    func SetSlice(slice []Slice) {
      anotherSlice = slice // "copy-boundary: copies a slice directly"
    }
    ```

* Chan Size

    ```golang
    c := make(chan int, 64) // "chan-size: channel size should be one or none"
    ```