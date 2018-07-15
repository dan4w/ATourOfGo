/**
 * Constants
 *  
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 7:46 PM
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import "fmt"

const Pi = 3.14

func main(){
	const World = "World of Warcraft"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)

}
