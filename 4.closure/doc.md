A closure in Go (and in many programming languages) is a function that "remembers" the variables from its surrounding scope even after the outer function has finished executing. This "remembered" state is encapsulated within the inner function, making it act like a self-contained unit.

## Example 1

```go
func createCounter() func() int {
    var counter int
    return func() int {
        counter++
        return counter
    }
}

func main() {
    counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
```

- The function `createCounter` defines a local variable `counter` initialized to `0`.  
- It returns an anonymous function (closure) that increments and returns the value of `counter`
- When `createCounter` is called, it returns a closure function.
- This closure captures the local variable counter within the `createCounter` scope.
- The variable `counter` in main holds the reference to this closure.
- Each call to `counter()` invokes the same closure.
- The closure has access to the `counter` variable captured when `createCounter` was called.
- On the first call, the closure increments `counter` from 0 to 1 and returns 1.
- On the second call, the same closure increments `counter` from 1 to 2 and returns 2.
- On the third call, it increments `counter` from 2 to 3 and returns 3.

### Why This Works

- **Closure**: When `createCounter` returns a closure, this closure "closes over" the `counter` variable. It means the closure retains a reference to the `counter` variable and can modify it.
- **Persistent State**: By storing the closure in the `counter` variable in main, we ensure that every call to `counter()` modifies the same `counter` variable, thus preserving its state between calls.
- **Separate Instances**: If `createCounter` is called multiple times, each call would create a new closure with its own separate counter variable. Therefore, calling `createCounter()` multiple times and invoking those closures would not share the state.

## Example 2

```go
	fmt.Println(createCounter()())
	fmt.Println(createCounter()())
	fmt.Println(createCounter()())
```

In this case, when you call `createCounter()()` in the `fmt.Println` statements, each invocation creates a new closure with its own separate `counter` variable. Therefore, each call increments a separate `counter` that starts from `0`, resulting in `1` each time.

## Example 3

```go
    myClosure := outer()
    myClosure() // Output: 10
```

In this example, `myClosure` is a function that has "captured" the variable `x` from the `outer` function's scope.


## Why Closures Matter in Go

- **Build "Objects"**: While Go doesn't have traditional classes, closures can mimic some object-oriented behaviors. *The `createCounter` and `createAdder` functions act like object constructors, and each returned closure is like an object instance with its own `counter` or `sum` variable*.

- **Hide Implementation Details**: The code using the closure doesn't need to know how the counting or adding is done internally.

- **Partial Application (Currying)** (see `5.` in the `snippets` folder):  Closures enable you to create functions that take some arguments now and others later. This is great for:

    - Creating Specialized Functions: The `createAdder` closure lets you make an adder function that's already "set up" with a starting value.
    - Improving Code Readability: You can break down complex function calls into simpler steps.

- **Callbacks and Event Handling**:  Closures are ideal for passing functions as arguments (callbacks) to other functions or for responding to events. This is because they can carry the necessary context along with the function logic.

- **Concurrency**: In Go, closures are essential for goroutines (lightweight threads). They allow you to safely share data and behavior between goroutines.

## Conclusion

The key to the expected behavior is that the closure retains access to the variables in its lexical scope (the environment where it was created). By storing and reusing the same closure instance, you ensure that the state (`counter` variable) is shared across multiple calls.