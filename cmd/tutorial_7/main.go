package main

import "fmt"

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nmemory location of thing1: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nresult is: %v", result)
	fmt.Printf("\nvalue of thing1 is: %v", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nmemory location of thing2: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

/*
func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("p points to: %v", *p)
	fmt.Printf("\nvalue of i is: %v", i)
	p = &i
	*p = 1
	fmt.Printf("\np points to: %v", *p)
	fmt.Printf("\nvalue of i is: %v", i)
	var k int32 = 2
	i = k
}
*/

/*
func main() {
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}
*/
