# Go By Example Tutorials

## `pointer.go`

```go
func main() {
	var p *int
	num := 12
	num1 := 33
	p = &num

	*p = 20

	p = &num1

	*p = 98

	fmt.Println(num, num1)
}
```
```go
func main() {
	var p1 *int // Declare a pointer to an int
	num1 := 12  // Initialize num1 with the value 12
	p1 = &num1  // p1 now holds the address of num1

	fmt.Println(p1) // Print the address stored in p1

	num2 := 34          // Initialize num2 with the value 34
	var p2 *int = &num2 // Declare p2 and assign it the address of num2

	fmt.Println(p2) // Print the address stored in p2

	p1 = p2         // Assign the address stored in p2 to p1
	fmt.Println(p1) // Print the new address stored in p1

	*p1 = 50          // Dereference p1 to change the value of num2 to 50
	fmt.Println(num2) // Print the new value of num2, prints: 50


// Output
    // 0xc00000a0b8
    // 0xc00000a0f0
    // 0xc00000a0f0
    // 50
}
```