# gouberfmt

Experimental golang static analysis tool that supports [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## Introduction

A static analysis is a function that inspects a package of Go code and reports a set of diagnostics (typically mistakes in the code), and perhaps produces other results as well, such as suggested refactorings or other facts.

The goal of this tool is to report missing idiomatic conventions in Go code based on [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## Installation

Will be added later

## List of Supported Analyzers

- Pointer to Interface

  ```golang
  // interface-pointer: var ifacePointer uses pointer to interface
  var ifacePointer *interface{}
  ```

- Pointer to sync.Mutex

  ```golang
  // "mutex-pointer: var lock uses pointer to sync.Mutex"
  var lock *sync.Mutex
  ```

- Copy Boundary

  ```golang
  func SetSlice(slice []Slice) {
    anotherSlice = slice // "copy-boundary: copies a slice directly"
  }
  ```

- Channel Size

  ```golang
  c := make(chan int, 64) // "chan-size: channel size should be one or none"
  ```

## Todo list

- [x] Pointers to Interfaces
- [ ] Receivers and Interfaces
- [x] Zero-value Mutexes
- [x] Copy Boundary
- [ ] Defer to Clean Up
- [x] Channel Size is One or None
- [ ] Start Enums at One
- [ ] Error Types
- [ ] Error Wrapping
- [ ] Don't Panic
- [ ] Prefer strconv over fmt
- [ ] Avoid string-to-byte conversion

## Contribution

Please read the [contributing guideline](https://github.com/gophersbd/gouberfmt/blob/master/contributing.md) guideline if you wish to contribute

## License

This projects is licensed under the [MIT License](https://github.com/gophersbd/gouberfmt/blob/master/LICENSE)
