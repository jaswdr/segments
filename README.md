# Segments

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
