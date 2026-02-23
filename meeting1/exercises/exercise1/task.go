package exercise1

//Write a function called Fibonacci accepting an int and returning n-th fibonacci number.
func Fibonacci(n int) int {
	a, b := 0, 1
	for range n {
		a, b = b, a+b
	}
	return a
}
