// Description: how private and public identifier work in go.
package main

import (
	"fmt"
	"math"
)

func main() {
	// Public identifier: Pi is exported from the math package
	fmt.Println("The value of Pi is approximately", math.Pi)
	// Private identifier: pi is not exported from the math package
	fmt.Println("The value of Pi is approximately", math.pi)
}

// Error: Due to the fact that you are trying to access Pi as private identifier. Change form pi to Pi