package main

import (
	"exportedApp/staff"
	"log"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 80000, Fulltime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 55000, Fulltime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, Fulltime: true},
	{FirstName: "Allan", LastName: "Andarson", Salary: 15000, Fulltime: false},
	{FirstName: "Elen", LastName: "Carter", Salary: 100000, Fulltime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	//log.Println(myStaff.All())
	staff.OverpaidLimit = 30000
	log.Println("Overpaid staff:", myStaff.Overpaid())
	log.Println("Underpaid staff:", myStaff.Underpaid())
}
