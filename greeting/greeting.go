package main

import "fmt"

func greeting(name string) string {
	return "Hi, " + name
}

func main() {
	fmt.Println(greeting("David"))
}
