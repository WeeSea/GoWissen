package main

import "fmt"

func main() {
	a := 10
	if b := a / 2; b > 5 {
		b := 10 / 2
		fmt.Println("b is bigger than 5", a, b)
	} else {
		fmt.Println("b is less than or equal to 5", a, b)
	}
	c := 3
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		if i == c {
			continue
		}
		fmt.Println(i)
	}

	//fmt.Println(i)

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	i = 0
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}

	}
	fmt.Println(i)

	s := "Hello, World!🤦‍♀️ 🌍🌍©®℗"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
