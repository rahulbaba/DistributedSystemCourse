package main

var c = make(chan int)
var a string

func main() {
	go func() {
		a = "hello, world\n"
		c <- 0
	}()

	<-c
	print(a)
}
