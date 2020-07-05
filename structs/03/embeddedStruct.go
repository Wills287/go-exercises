package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	s := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "Blue",
		},
		luxury: true,
	}
	t := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "White",
		},
		fourWheel: false,
	}

	fmt.Println(s)
	fmt.Println(t)
}
