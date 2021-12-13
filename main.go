// ============================================================================
// = main.go																  =
// = 	Description: used to run the goconvert package functions			  =
// = 	Date: December 12, 2021												  =
// ============================================================================

package main

import (
	"fmt"
	n2w "goconvert/num2words"
	"goconvert/utils"
)

func main() {
	num := n2w.Pow(n2w.New(10), n2w.New(606))
	utils.PrintNumber(num)
	fmt.Println(n2w.Num2Words(num, true))
}