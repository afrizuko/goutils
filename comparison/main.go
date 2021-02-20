package main

import "fmt"

func main() {

	name := "Jesus Is Lord"

	fmt.Println(string(name[0]))
	fmt.Println(string([]rune(name)[0]))
}
