package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(11))
	isPalindrome(121)
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var r int
	for x > r {
		r = r*10 + x%10
		x /= 10
	}
	return x == r || x == r/10
}
