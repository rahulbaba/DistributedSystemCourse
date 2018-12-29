package main

var a string

func main() {
	go func() { a = "hello" }()
	print(a)
}
