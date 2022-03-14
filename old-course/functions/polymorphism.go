package main

import "fmt"

type Engineer interface {
	MakeStuff() string
}

type CSEngg struct {
	Laptop bool
	Css    bool
	Html   bool
	WebDev bool
	Golang bool
}

func (e *CSEngg) MakeStuff() string {
	tc := 0
	x := "\nSoftware Engineer made stuff with\n"
	if e.Laptop {
		tc++
		x += fmt.Sprint(tc) + ". Laptop\n"
	}
	if e.Css {
		tc++
		x += fmt.Sprint(tc) + ". CSS\n"
	}
	if e.Html {
		tc++
		x += fmt.Sprint(tc) + ". HTML\n"
	}
	if e.WebDev {
		tc++
		x += fmt.Sprint(tc) + ". Web Development\n"
	}
	if e.Golang {
		tc++
		x += fmt.Sprint(tc) + ". Golang\n"
	}
	return x
}

type MechEngg struct {
	Helmet     bool
	Overalls   bool
	Blueprints bool
}

func (e *MechEngg) MakeStuff() string {
	tc := 0
	x := "\nMechanical Engineer made stuff with\n"
	if e.Helmet {
		tc++
		x += fmt.Sprint(tc) + ". Helmet\n"
	}
	if e.Overalls {
		tc++
		x += fmt.Sprint(tc) + ". Overalls\n"
	}
	if e.Blueprints {
		tc++
		x += fmt.Sprint(tc) + ". Blueprints\n"
	}
	return x
}

type ECEngg struct {
	Gloves         bool
	CircuitDiagram bool
	BreadBox       bool
	LEDs           bool
}

func (e *ECEngg) MakeStuff() string {
	tc := 0
	x := "\nElectrical Engineer made stuff with\n"
	if e.Gloves {
		tc++
		x += fmt.Sprint(tc) + ". Gloves\n"
	}
	if e.CircuitDiagram {
		tc++
		x += fmt.Sprint(tc) + ". Circuit Diagram\n"
	}
	if e.BreadBox {
		tc++
		x += fmt.Sprint(tc) + ". Breadbox\n"
	}
	if e.LEDs {
		tc++
		x += fmt.Sprint(tc) + ". LEDs\n"
	}
	return x
}

func foo(e Engineer) {
	fmt.Println(e.MakeStuff())
}

func bar(e Engineer) {
	switch e.(type) {
	case *MechEngg:
		fmt.Printf("\n---------------------\nMechanical Engineer\n---------------------\nHas Helmet: %t\n", e.(*MechEngg).Helmet)
	case *CSEngg:
		fmt.Printf("\n---------------------\nComputer Science Engineer\n---------------------\nKnows CSS: %t\n", e.(*CSEngg).Css)
	case *ECEngg:
		fmt.Printf("\n---------------------\nElectrical Engineer\n---------------------\nHas Breadbox: %t\n", e.(*ECEngg).BreadBox)

	}
}

func main() {
	cse := CSEngg{Laptop: true, Css: false, Html: true, WebDev: true, Golang: true}
	fmt.Println(cse)

	me := MechEngg{
		Helmet:     true,
		Overalls:   true,
		Blueprints: true,
	}
	fmt.Println(me)

	ece := ECEngg{
		Gloves:         true,
		CircuitDiagram: true,
		BreadBox:       true,
		LEDs:           true,
	}
	fmt.Println(ece)

	bar(&cse)
	foo(&cse)

	bar(&me)
	foo(&me)

	bar(&ece)
	foo(&ece)
}
