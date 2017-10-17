// This solution for the fifth Project Euler problem was written by Evert "ElecProg" Provoost
// Please try to find a solution yourself before taking a look at this!

package main

func main() {

	/*
		We'll only check if a number is divisible by:
		* 20 (includes checks for 10, 5, 4 and 2)
		* 19
		* 18 (includes checks for 9, 6, 3 and 2)
		* 17
		* 16 (includes checks for 8, 4 and 2)
		* 15 (includes checks for 5 and 3)
		* 14 (includes checks for 7 and 2)
		* 13
		* 12 (includes checks for 6, 4, 3 and 2), or
		* 11
		All numbers smaller than 11 are already checked for.
	*/

	dividers := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// The problem states that this is the first number divisible by 1..10
	// So any number before this number isn't divisble by 1..20
	num := 2520

	// This loop keeps running until the number is divisible
	divisible := false
	for !divisible {
		// We only have to check even nubers, else it isn't divisible by 2.
		num += 2

		// Check every divider
		divisible = true
		for _, div := range dividers {
			if num%div != 0 {
				divisible = false
				break
			}
		}
	}

	print("The solution is ")
	println(num)
}
