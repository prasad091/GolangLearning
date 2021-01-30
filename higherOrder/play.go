package main

import "fmt"

const message = "Hello Prasad Thangavel"

func main() {
	g := greeter{
		greeting: "hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println(message)

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
