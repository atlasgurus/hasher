
# README for the `hasher` Package

## Overview
The `hasher` package, part of the `github.com/atlasgurus/hasher` module, provides a utility for generating SHA-256 hashes of Go structs. This package is useful for creating consistent, unique representations of struct instances, which can be used in applications like caching, comparison, or as part of complex data structures.

## Installation
To use the `hashier` package in your Go project, import it as follows:

```go
import "github.com/atlasgurus/hasher"
```

Ensure that the package is correctly located in your project's workspace or include it as a dependency in your `go.mod` file.

## Usage
The primary function in the `hasher` package is `ComputeHash`, which takes any Go struct as input and returns a SHA-256 hash of the struct as a byte slice.

```go
func ComputeHash(s interface{}) []byte
```

## Features
- **Flexible Struct Hashing:** The `ComputeHash` function can handle various types of struct fields, including nested structs.
- **Selective Field Hashing:** Fields in structs can be excluded from the hash computation using a struct tag `hash:"-"`.
- **Support for Basic Data Types:** Handles basic Go data types like integers, strings, and booleans.

## Example
Here's a simple example of how to use the `hasher` package:

```go
package main

import (
    "fmt"
    "github.com/atlasgurus/hasher"
)

type Person struct {
    Name   string
    Age    int `hash:"-"`
    Parent *Person
}

func main() {
    parent := &Person{Name: "Jane Doe", Age: 50}
    person := &Person{Name: "John Doe", Age: 30, Parent: parent}

    parentHash := hasher.ComputeHash(parent)
    personHash := hasher.ComputeHash(person)

    fmt.Printf("Parent Hash: %x\n", parentHash)
    fmt.Printf("Person Hash: %x\n", personHash)
}
```

## Testing
To test the functionality of `ComputeHash`, the package includes a test case `TestComputeHash` in `hasher_test.go`. It verifies the hash output for predefined struct instances.

## Contribution
Contributions to the `hasher` package are welcome. Please ensure that any pull requests include relevant tests and adhere to Go's coding standards.

## License
MIT License