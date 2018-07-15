/**
 * Page 5/17
 * Simplified arguements of functions
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14T16:06:32-07:00
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:55:28-07:00
 */



package main


import "fmt"

// instead of x int, y int we do x, y int
func add(x,y int) int {
	return x + y
}

func main(){
	fmt.Println(add(42,13))
}
