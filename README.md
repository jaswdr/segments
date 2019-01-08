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
    s := segments.Start()
    // make some expensive operation
    s.Stop()

    fmt.Printf("Elapse: %v", s.Diff())
}
```

## Author

Jonathan A. Schweder <jonathanschweder@gmail.com>
