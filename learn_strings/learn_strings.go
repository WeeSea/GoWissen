package main

import "fmt"

func main() {
	var s string
	s = "Hello, World!"
	fmt.Println(s)

	t := "Hello, World of T!"
	fmt.Println(t)

	u := "Hello, \n\t\"World!\" with a backslash \\"
	fmt.Println(u)

	v := `Hello,
	"World!" with a backslash \`
	fmt.Println(v)

	w := "Hello, world!"
	w2 := "👋 🌍"
	w3 := w + " " + w2
	fmt.Println(w3)

	x := w[0]
	x2 := w[4]
	fmt.Println(w, x, string(x), x2, string(x2))

	s2 := s[0:5]
	s3 := s[7:12]
	s4 := s[:5]
	fmt.Println(s4)
	s5 := s[7:]
	fmt.Println(s5)
	fmt.Println(len(s), s, s2, s3, s4, s5, len(s5))
	w4 := "👋 🌍"
	w5 := w2[:1]
	w6 := w2[2:]
	fmt.Println(w4, len(w4), w5, len(w5), w6, len(w6))
	z := "Hello, "
	//var r rune = 127757
	r := '🌍'
	z = z + string(r)
	fmt.Println(z)

	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals, vals[0], vals[1], vals[2])
}
