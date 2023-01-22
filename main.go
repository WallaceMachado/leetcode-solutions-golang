package main

import "fmt"

func main() {
	teste := []string{"flower", "flow", "flight"}
	r := longestCommonPrefix(teste)
	fmt.Println(r)
}

func longestCommonPrefix(strs []string) string {
	r := ""
	c := 0
	x := 0
	var a uint8
	for i := 0; i <= c; i++ {
		fmt.Println("i: ", i)
		a = 0
		x = 0
		for _, v := range strs {
			fmt.Println("ci: ", c)
			if c == 0 {
				c = len(v) - 1
			} else if len(v) < c {
				c = len(v) - 1
			}
			fmt.Println("cf: ", c)

			fmt.Println("a: ", a)
			fmt.Println("vi: ", v[i])
			fmt.Println("xi: ", x)
			if a == 0 {
				a = v[i]
				x++
				fmt.Println("a0: ", a)
			} else if a == v[i] {
				x++
			}
			fmt.Println("xf: ", x)

		}
		fmt.Println("ri: ", r)
		if x == len(strs) {
			r += string(a)
		}
		fmt.Println("rf: ", r)
	}

	return r

}
