package main

import (
	"fmt"
)

//multi fx
func vals() (int, int) {
	return 3, 7
}

//variadic fx
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//anonymous functions and closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//Testgo to test go capabilities
func Testgo() {
	//arrays
	// var a [5]int
	// fmt.Println("emp:", a)
	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	//     for j := 0; j < 3; j++ {
	//         twoD[i][j] = i + j
	//     }
	// }
	// fmt.Println("2d: ", twoD)

	//slices
	// s := make([]string, 6)
	// fmt.Println("emp:", s)
	// l := s[2:5]
	// fmt.Println("sl1:", l)
	// t := []string{"g", "h", "i"}
	// fmt.Println("dcl:", t)

	//maps
	// m := make(map[string]int)
	// m["k1"] = 7
	// m["k2"] = 13
	// fmt.Println("map:", m)

	//ranges
	// nums := []int{2, 3, 4}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("sum:", sum)
	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }

	//multi return fxs
	// a, b := vals()
	// fmt.Println(a)
	// fmt.Println(b)

	// _, c := vals()
	// fmt.Println(c)

	//variadic fxs
	// sum(1, 2)
	// sum(1, 2, 3)
	// nums := []int{1, 2, 3, 4}
	// sum(nums...)

	//anonymous functions and closures
	// nextInt := intSeq()
	// // See the effect of the closure by calling nextInt a few times.
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// newInts := intSeq()
	// fmt.Println(newInts())

	//Recursion
	fmt.Println(fact(7))
}
