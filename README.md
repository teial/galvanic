# My Go Library

Welcome to Galvanic! This library is built using Go 1.23rc1 and takes advantage of the new range function iterators feature. It is designed to help Go developers create a compose range iterators for complex data processing in a functional style.

## Features

- **Range Function Iterators**: Leverage the new iterator functionality in Go 1.23rc1.
- **Easy Integration**: Simple and straightforward to use in your existing Go projects.

## Usage

Here is a basic example of how to use My Go Library:

```go
package main

import (
    "fmt"
    "github.com/teial/galvanic/sequence"
)

func main() {
    // Example of using the range function iterators. This code outputs 5 3
	iter := sequence.Values(1,4,2,3,5).Reverse().Take(3).Filter(func(e int) bool { return e%2 == 1 })
    for val := range iter {
        fmt.Println(val)
    }
}
```

Due to the way generic in Go work, there are some limitations to what is achievable. In particular, it is not possible to have chains that convert types. which means `Map` should be a function and not a method. While it might sound limiting, it actually makes resulting code closer following Go philosophy. By splitting iterator pipelines into several chunks, you will have easieer time reading and maintaining such code.

## Contributing

We welcome contributions to Galvanic! If you find a bug or have a feature request, please open an issue on GitHub. If you would like to contribute code, please fork the repository and create a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

Special thanks to the Go development team for the new features in Go 1.23rc1 and to all contributors for their invaluable input and support.
