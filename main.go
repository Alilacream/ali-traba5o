package main

import "fmt"
import "ali/ali"

func main() {
	var nb int64 = 16
	var pt *int64 = &nb
	fmt.Printf("the sqrt of nb is equal to %d\n", ali.Sqrt(*pt))
}
