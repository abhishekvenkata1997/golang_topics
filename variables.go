package main

import (
	"fmt"
	"strconv"
)

var actorName string = "Abhishek"
var companion string = "venkata"

// or
var (
	doctorNumv int = 11
	season     int = 5
	i          int = 15
)

func main() {
	//inner scope can be declared again
	//need to use variables if they are declared
	var i int
	i = 42
	//or
	var j int = 17
	//or
	k := 25
	var s float32 = 25
	var str string
	//str = string(i)
	str = strconv.Itoa(i)
	fmt.Println(i + j + k)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T", str, str)
}
