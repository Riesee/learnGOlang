package main

import (
	"fmt"
	"time"
)

var x, y int = 3, 2
func main() {
	var i int = 42
	var f float64 = float64(i) + 0.8
	var u uint = uint(f)

	fmt.Printf("uuu %T  \n ", i )
	fmt.Printf("uuu %T  \n ", f )
	fmt.Printf("uuu %T  \n ", u )
	fmt.Println("time:", time.Now())
}