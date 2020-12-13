package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	a := 0
	b := 1
	fmt.Println(b / a)
	fmt.Println("Am I dead?")
}
