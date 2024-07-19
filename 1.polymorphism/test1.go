// Type Assertions using interface.

package main

import "fmt"

func processValue(x interface{}) {
    // Value assertion with type check
    if i, ok := x.(int); ok {
        fmt.Println("x is an int:", i)
    } else {
        fmt.Println("x is not an int")
    }

    y := x.(float64) // Will panic if x is not a float64 type.
    fmt.Println(y)
    
    // Type switch
    switch v := x.(type) {
    case int:
        fmt.Println("x is an int:", v)
    case string:
        fmt.Println("x is a string:", v)
    case float64:
        fmt.Println("x is a float64:", v)
    case bool:
        fmt.Println("x is a bool:", v)
    default:
        fmt.Println("x is of unknown type")
    }
}

func Test1() {
    // Test with various data types
    processValue(42)             // int
    processValue("Hello, Go!")    // string
    processValue(3.14159)       // float64
    processValue(true)           // bool
    processValue([]int{1, 2, 3}) // slice (unknown type)
}