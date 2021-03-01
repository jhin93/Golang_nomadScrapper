package main

import "fmt"

// Arrays and Slices
// Go에서는 Array를 만들 때 길이를 명시해주어야 한다. 또한 어떤 타입의 Array인지도 명시해주어야 한다.

func main() {
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "lalala"
	names[4] = "jhin"
	fmt.Println(names)
}
