package main

import (
	"fmt"
)

func main() {
	/* Example 1 */
	// fmt.Println("Hello World!")

	/* Example 2 */
	// var favNums = [5]float64{1, 2, 3, 4, 5}

	// for key, value := range favNums {
	// 	fmt.Println(value, key)
	// }

	/* Example 2 */
	// We can store multiple items in a map as well

	// superhero := map[string]map[string]string{
	// 	"Superman": map[string]string{
	// 		"realname": "Clark Kent",
	// 		"city":     "Metropolis",
	// 	},

	// 	"Batman": map[string]string{
	// 		"realname": "Bruce Wayne",
	// 		"city":     "Gotham City",
	// 	},
	// }

	// We can output data where the key matches Superman

	// if hero, flag := superhero["Superman"]; flag {
	// 	fmt.Println(hero["realname"], hero["city"], flag)
	// } else {
	// 	fmt.Println("GG")
	// }

	/* Example 3 */
	// defer printTwo()
	// printOne()

	/* Example 4 */
	// fmt.Println("3 / 0 =", safeDiv(3, 0))
	// fmt.Println("4 / 2 =", safeDiv(4, 2))

	// demPanic()

	/* Example 4 */
	// fmt.Println("What is your name? ")
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
	// fmt.Println("Hello, ", text)
}

// Used to demonstrate defer

func printOne() { fmt.Println(1) }

func printTwo() { fmt.Println(2) }

// If an error occurs we can catch the error with recover and allow
// code to continue to execute

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

// Demonstrate how to call panic and handle it with recover

func demPanic() {

	defer func() {

		// If I didn't print the message nothing would show

		fmt.Println(recover())

	}()
	panic("PANIC")

}
