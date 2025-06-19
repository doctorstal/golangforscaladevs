package main

import "fmt"

func IfsAndLoopsExample() {
	a, b := 1, 2
	if a == b {
		panic("1 is not equal 2!")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Next step:", i)
	}

	// You can ommit parts of the loop statement:
	for i := 0; i < 5; {
		i++
	}

	i := 0
	for ; i < 5; i++ {
	}

	for ; ; i++ {
		if i > 5 {
			break
		}
	}

	for i < 5 {
		i++
	}

	for {
		fmt.Println("loop")
		break
	}

	for i := range 5 {
		fmt.Println("Next range step:", i)
	}

	slice := make([]int, 5)

	for i := range slice {
		slice[i] = i
	}

	slice = append(slice, 9, 10)

	for _, v := range slice {
		fmt.Println("element ", v)
	}
}
