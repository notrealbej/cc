package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//print
	fmt.Println("hello world")

	//variables
	var a = "This is string"
	var c, d int = 1, 2

	fmt.Println(a)
	fmt.Println(c + d)

	//loops
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	//conditional
	if c%2 == 0 {
		fmt.Println("wtf")
	} else {
		fmt.Println("we good")
	}

	//arrays
	var arr [5]int
	fmt.Println(arr)
	arr[1] = 20
	fmt.Println(arr)
	fmt.Println(arr[1])

	//fuctions and user input
	var x, y int
	fmt.Print("Enter number: ")
	fmt.Scan(&x, &y)
	fmt.Println(add(x, y))

	//Service
	// http.HandleFunc("/", helloHandler)
	// http.ListenAndServe(":10000", nil)

	//go routine
	go printMessage()
	fmt.Println("Main func")
	time.Sleep(time.Second)
}

func printMessage() {
	time.Sleep(1)
	fmt.Println("Goroutine!")
}

// function
func add(a int, b int) int {
	return a + b
}

// service func
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
