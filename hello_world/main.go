package main

import "fmt"

func Hello(who string) string {
	if who == "" {
		who = "World"
	}
	return "Hello, " + who
}

func main() {
	fmt.Println(Hello("World"))
}

func Repeat(character string) string {
	var repeatCount = 5
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
