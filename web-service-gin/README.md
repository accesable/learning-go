# Gin (Golang RestAPI framework)

the mechanism of Gin work alike to `Express.js` in NodeJS. First we initilize an server instance from the framework and pass callbacks as handler to handle the incomming requests.\
In Go, which is a statically-typed language, callback functions must be defined as a specific type in order to be used as parameters to other functions or methods. This ensures type safety and allows the compiler to enforce type correctness.

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    // Route handler (callback function)
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    router.Run(":8080")
}
```

```js
const express = require('express');
const app = express();

// Route handler (callback function)
app.get('/hello', (req, res) => {
    res.json({ message: 'Hello, World!' });
});

app.listen(8080, () => {
    console.log('Server is running on port 8080');
});
```

## The `:=` Operator

The `:=` symbol in Go is called the short variable declaration operator. It is used to declare and initialize a new variable. It combines variable declaration and assignment into a single statement, making the code more concise and readable.

When you use `:=`, Go infers the variable's type based on the value on the right-hand side of the expression. This is called type inference. It's particularly useful when the type of the variable is obvious from the initialization value, as it reduces repetition and makes the code more succinct.

Example
```go
// using :=
id := c.Param("id")
// using =
var id string
id = c.Param("id")


// using =
var p *int    // Declaration of a pointer variable p of type *int
var x int = 10
p = &x       // Assigning the memory address of x to p (address-of operator &)
// using :=
x := 10     // Declare and initialize a variable x with the value 10
p := &x     // Declare and initialize a pointer variable p with the memory address of x
```

## Loop in Go

In Go, loops are constructs used to repeatedly execute a block of code until a specified condition is met. There are three types of loops in Go: the `for` loop, the for `range` loop, and the `while` loop.\
The for `range` loop is used to iterate over elements of a slice, array, string, map, or channel. It iterates over each element, assigning both the index and value to variables (or using _ to ignore one or both).\
Go doesn't have a built-in `while` loop, but you can achieve the same effect using a for loop with only a condition.

```go
// for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}


for index, element := range collection {
    // Code to be executed
}
// for range loop
numbers := []int{1, 2, 3, 4, 5}
for index, num := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, num)
}

// While loop
x := 0
for x < 5 {
    fmt.Println(x)
    x++
}

// Infinite Loop
for {
    fmt.Println("This loop will run indefinitely")
}
```

## Pointer in Go

Pointers in Go, like in many other programming languages, are variables that store the memory address of another variable. They allow indirect access to the value or memory location of a variable. Pointers are widely used in Go for various purposes, including passing parameters by reference, dynamic memory allocation, and implementing data structures.
**Declaration and Initialization:**
You declare a pointer variable using the * symbol followed by the type of the variable it points to:

```go
var ptr *int   // Declaration of a pointer variable ptr of type *int
```

You can initialize a pointer variable by assigning the memory address of another variable to it using the address-of operator `&`:

```go
var x int = 10
ptr = &x       // Assigning the memory address of x to ptr
```

**Dereferencing**

```go
fmt.Println(*ptr)   // Prints the value stored at the memory address pointed to by ptr
```

**Zero Value of Pointers:**
If a pointer is declared but not initialized, its zero value is `nil`, indicating that it doesn't point to any valid memory address.

**Pass by Reference:**\
In Go, function parameters are passed by value by default, meaning a copy of the value is passed to the function. However, you can pass a pointer to a value if you want to pass it by reference:

```go
package main

import "fmt"

// swap swaps the values of two integers using pointers.
func swap(x, y *int) {
    // Dereference the pointers and swap the values
    *x, *y = *y, *x
}

func main() {
    // Declare two integers
    a, b := 10, 20

    // Print the initial values
    fmt.Println("Before swapping:")
    fmt.Println("a:", a)
    fmt.Println("b:", b)

    // Call the swap function, passing the addresses of a and b
    swap(&a, &b)

    // Print the swapped values
    fmt.Println("\nAfter swapping:")
    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```
**Use Cases:**

- Dynamic Memory Allocation: Allocating memory dynamically using new() or make() functions.
- Passing Parameters by Reference: Modifying the original value of a variable inside a function.
- Efficient Data Structures: Implementing linked lists, trees, and other data structures.

Pointers in Go are powerful but require careful handling to avoid common pitfalls like nil pointer dereferences and memory leaks. However, they provide flexibility and efficiency in managing memory and accessing data.

--

Understanding *Dereferencing* : Dereferencing a pointer means accessing the value that is stored at the memory address pointed to by the pointer.

In Go, when you define a method on a type (a struct type, for example) and use a pointer receiver, Go automatically handles the dereferencing for you when you call the method on a pointer to that type.
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Method with a pointer receiver
func (p *Person) ChangeName(newName string) {
    p.Name = newName
}

func main() {
    // Create a Person struct instance
    person := Person{Name: "Alice", Age: 30}

    // Create a pointer to the Person struct instance
    ptr := &person

    // Call the method using the pointer
    ptr.ChangeName("Bob")

    // Print the updated name
    fmt.Println("Name:", person.Name) // Output: Name: Bob
}
```

In C++, the behavior is similar. When you define a member function (method) in a class or struct and use a pointer as its parameter, you can call that function using a pointer to an object, and the compiler will automatically dereference the pointer for you.

```cpp
#include <iostream>
using namespace std;

class Person {
public:
    string name;
    int age;

    // Method with a pointer parameter
    void changeName(const string& newName) {
        name = newName;
    }
};

int main() {
    // Create a Person object
    Person person{"Alice", 30};

    // Create a pointer to the Person object
    Person* ptr = &person;

    // Call the method using the pointer
    ptr->changeName("Bob");

    // Print the updated name
    cout << "Name: " << person.name << endl; // Output: Name: Bob

    return 0;
}
```
