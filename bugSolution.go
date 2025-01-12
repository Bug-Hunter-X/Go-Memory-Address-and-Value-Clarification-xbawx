```go
package main

import "fmt"

func main() {
    var x int = 10
    fmt.Println(x) // Print the value of x
    fmt.Printf("Memory Address of x: %p\n", &x) // Print the memory address of x
    ptr := &x
    fmt.Printf("Value pointed to by ptr: %d\n", *ptr) // Print the value at the memory address stored in ptr
}
```