package main

import (
	"fmt"
)

// Base struct with default implementations
type baseListener struct {}

// Default method implementations at base level
func (bl *baseListener) buttonDownListener(origin string) {
	fmt.Println("BaseListener buttonDownListener")
}

func (bl *baseListener) buttonUpListener(origin string) {
	fmt.Println("BaseListener buttonUpListener")
}

// Concrete listener extending the base one
type concreteListener struct {
	*baseListener
}

// The method buttonDownListener is overrider by the concrete listener
func (cl *concreteListener) buttonDownListener(origin string) {
	fmt.Println("ConcreteListener buttonDownListener")
}


func main() {
	var listener concreteListener

	listener.buttonDownListener("buttonDownEvent")
	listener.buttonUpListener("buttonUpEvent")
}
