package main

type rectiint struct {
	width, height int
}

// This area method has a receiver type of *rectiint.

func (r *rectiint) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
//  Here’s an example of a value receiver.
func (r rectiint) perim() int {
	return 2*r.width + 2*r.height
}
