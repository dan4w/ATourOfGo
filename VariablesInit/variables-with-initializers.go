/**
 * VariablesInit
 * Page 9/17
 * A var declaration can include initializers, one per variables.
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 5:19 PM ----2018-07-14T02:59:28-07:00
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import "fmt"

var i, j int = 1,2

func main(){
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
