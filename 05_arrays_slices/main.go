package main

import (
	"fmt"
)

func main() {
	// Arrays (fixed values)
	arr := [2]string{"apple", "banana"}
	fmt.Println(arr, arr[1])

	//Slice
	slc := []string{"apple", "banana"}
	slc = append(slc, "orange")
	fmt.Println(slc, len(slc), slc[:2], slc[2:])
}
