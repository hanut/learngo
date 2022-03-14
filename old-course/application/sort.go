package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	Name   string
	Age    uint8
	Salary int32
}

type ByAge []Person
type ByName []Person
type BySalary []Person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }

func (p ByName) Len() int           { return len(p) }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }

func (p BySalary) Len() int           { return len(p) }
func (p BySalary) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p BySalary) Less(i, j int) bool { return p[i].Salary < p[j].Salary }

func (p Person) String() string {
	return fmt.Sprintf("Name: %s\tAge: %d\tSalary: %d\n", p.Name, p.Age, p.Salary)
}

func main() {
	normalSort()
	customSort()
}

func normalSort() {
	x := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"Hanut", "Aditya", "Simran", "Gurleen", "Ritu", "Amit"}
	fmt.Println("--------------------------------\nBefore Sort\n--------------------------------")
	fmt.Println(x)
	fmt.Println(xs)

	sort.Ints(x)
	sort.Strings(xs)

	fmt.Println("\n--------------------------------\nAfter Sort\n--------------------------------")
	fmt.Println(x)
	fmt.Println(xs)
}

func customSort() {
	fmt.Println("\n\n--------------------------------\nCustom Sorts\n--------------------------------")
	people := []Person{
		{Name: "Vikas", Age: 22, Salary: 1600000},
		{Name: "Wazim", Age: 23, Salary: 900000},
		{Name: "Arun", Age: 40, Salary: 500000},
		{Name: "Grover", Age: 21, Salary: 3200000},
		{Name: "Jamshed", Age: 32, Salary: 7000000},
	}

	bigList := make([]Person, 10000, 10000)
	for i := range bigList {
		switch {
		case (i%5 == 0):
			bigList[i] = people[4]
		case (i%4 == 0):
			bigList[i] = people[3]
		case (i%3 == 0):
			bigList[i] = people[2]
		case (i%2 == 0):
			bigList[i] = people[1]
		default:
			bigList[i] = people[0]
		}
	}

	fmt.Println("Before sort:")
	fmt.Print(people)

	sort.Sort(ByAge(people))
	fmt.Println("\n\nSorted By Age")
	fmt.Print(people)

	sort.Sort(ByName(people))
	fmt.Println("\n\nSorted By Name")
	fmt.Print(people)

	sort.Sort(BySalary(people))
	fmt.Println("\n\nSorted By Salary")
	fmt.Print(people)
	sort.Sort(sort.Reverse(BySalary(people)))
	fmt.Println("\n\nSorted (REVERSE) By Salary")
	fmt.Print(people, "\n")

	fmt.Println("sorting bigList[", len(bigList), "] by age")
	sat := time.Now()
	sort.Sort(ByAge(bigList))
	fmt.Println("Time taken to sort:", time.Since(sat))
}
