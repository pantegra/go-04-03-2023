package main

import (
	"fmt"
)

type astro struct {
	name string
	age int
	mission string
	isneeded bool
}

func main() {
	astro1 := astro{"Megan",39,"Mars", false}
	astro2 := astro{"Trump",88,"Moon",false}
	astro3 := astro{"Biden",90,"Alpha Centare",false}

	fmt.Println(astro1)
	fmt.Println(astro2)
	fmt.Println(astro3)

	p := []astro{astro1,astro2,astro3}

	fmt.Println(p)

}