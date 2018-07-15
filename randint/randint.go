/**
 * Page 1/17
 * This does not generate a seed for a random number.
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14T15:50:41-07:00
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:58-07:00
 */



package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
