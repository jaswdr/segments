# Segments 

[![Build Status](https://travis-ci.com/jaswdr/segments.svg?branch=master)](https://travis-ci.com/jaswdr/segments)
[![codecov](https://codecov.io/gh/jaswdr/segments/branch/master/graph/badge.svg)](https://codecov.io/gh/jaswdr/segments)
[![Maintainability](https://api.codeclimate.com/v1/badges/de70193c532a883b77ac/maintainability)](https://codeclimate.com/github/jaswdr/segments/maintainability)

> Monitor and record segments of expensive code

### Quick example

```go
package main

import (
    "fmt"

    "github.com/jaswdr/segments"
)

func main() {
    s1 := segments.Start()
    // expensive operation
    s1.Stop()

    fmt.Printf("Elapse: %v", s1.Diff())

    // Wrapping a function
    s2 := segments.Wrap(func() {
        // expensive operation
    })

    fmt.Printf("Elapse: %v", s2.Diff())
}
```

## Author

Jonathan A. Schweder <jonathanschweder@gmail.com>
