```go
package main

import "fmt"

func main() {
    var x int
    fmt.Println(x) // This will print 0
    x++
    fmt.Println(x) // This will print 1
    fmt.Println(&x) // This will print the memory address of x
    fmt.Println(*(&x)) //This will print the value at memory address of x, which is 1 
}
```