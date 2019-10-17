# gouberfmt

Experimental golang static analysis tool that supports - https://github.com/uber-go/guide/blob/master/style.md

## Supported Analyzers

* Pointer to Interface

    ```
    // interface-pointer: var ifacePointer uses pointer to interface
    var ifacePointer *interface{}
    ```
* Pointer to sync.Mutex

    ```
    // "mutex-pointer: var lock uses pointer to sync.Mutex"
    var lock *sync.Mutex
    ```
