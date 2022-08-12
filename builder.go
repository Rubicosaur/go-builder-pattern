package main

import "fmt"

// builder pattern

type human struct {
	age       int
	height    float64
	eyeColour string
}

func (h human) withEyeColour(coluor string) human {
	h.eyeColour = coluor

	return h
}

func (h human) withAge(age int) human {
	h.age = age

	return h
}

func (h human) withHeight(height float64) human {
	h.height = height

	return h
}

func newHuman() human {
	return human{}
}

func (h human) reset() human {
	h = newHuman()
	return h
}

func newHumanWithFields(age int, height float64, eyeColour string) human {
	return human{
		age:       age,
		height:    height,
		eyeColour: eyeColour,
	}
}

func giant() human {
	return newHuman().withHeight(255.55).withEyeColour("green")
}

func main() {
	me := newHuman().withEyeColour("green").withAge(20).withHeight(177.88)

	fmt.Println(me)

	you := newHumanWithFields(22, 166.66, "blue")

	fmt.Println(you)

	youngGiant := giant().withAge(3)
	oldGiant := giant().withAge(199)

	fmt.Println(youngGiant, oldGiant)
}
