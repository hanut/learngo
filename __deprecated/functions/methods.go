package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       uint8
	Vehicles  []Vehicle
}

type Vehicle struct {
	Model   string
	Make    string
	Year    string
	Color   string
	Variant string
}

func (v *Vehicle) String() string {
	return fmt.Sprintf(`---------------------------------------------
Make:		%s
Model:		%s
Variant:	%s
Year:		%s
Color:		%s
---------------------------------------------`, v.Make, v.Model, v.Variant, v.Year, v.Color)
}

func (p *Person) String() string {
	desc := fmt.Sprintf(`---------------------------------------------
First Name:	%s
Last Name:	%s
Age:		%d,
---------------------------------------------`, p.FirstName, p.LastName, p.Age)
	desc += "\nVehicles"
	for _, v := range p.Vehicles {
		desc += "\n" + v.String()
	}
	return desc
}

func main() {
	p1 := Person{
		FirstName: "Hanut",
		LastName:  "Singh Gusain",
		Age:       33,
		Vehicles: []Vehicle{
			{
				Model:   "Model X",
				Make:    "Tesla",
				Year:    "2022",
				Color:   "White",
				Variant: "Snazzy",
			},
			{
				Model:   "Hector",
				Make:    "MG",
				Year:    "2021",
				Color:   "Matte Black",
				Variant: "2.0 Diesel MT0",
			},
		},
	}

	fmt.Println(p1.String())
}
