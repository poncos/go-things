package main

import (
	"fmt"
)

type baseListener struct {}

func (bl *baseListener) buttonDownListener(origin string) {
	fmt.Println("BaseListener buttonDownListener")
}

func (bl *baseListener) buttonUpListener(origin string) {
	fmt.Println("BaseListener buttonUpListener")
}

type concreteListener struct {
	*baseListener
}

func (cl *concreteListener) buttonDownListener(origin string) {
	fmt.Println("ConcreteListener buttonDownListener")
}


func main() {
	var listener concreteListener

	listener.buttonDownListener("buttonDownEvent")
	listener.buttonUpListener("buttonUpEvent")
}
