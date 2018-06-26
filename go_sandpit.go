package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hola Mundo\n")
	// fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	// Testgo()
	// Testpo()
	// Testst()
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	rp := &r
	fmt.Println("area defered: ", rp.area())
	c := circle{radius: 5}
	measure(r)
    measure(c)
}
