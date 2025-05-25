// Description: This Go program imports the "math/rand" package and uses it to generate a random number.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My lucky number usually is ", rand.Intn(10))
}
