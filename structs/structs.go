package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	fullTime bool
}

// Data Struct with Anonymous fields
type Data struct {
	string
	int
	bool
}

type Salary struct {
	fixedPay     int
	joiningBonus int
	variablePay  int
}

// EmployeeData Nested struct
type EmployeeData struct {
	name     string
	salary   Salary
	fullTime bool
}

// EmployeeInfo Promoted fields, Here fields of Salary struct can be directly accessed, called Promoted fields
type EmployeeInfo struct {
	name string
	Salary
	fullTime bool
}

// Nested interface
type Salaried interface {
	getSalary() int
}

type EmployeeSalary struct {
	fixedPay     int
	joiningBonus int
	variablePay  int
}

func (es EmployeeSalary) getSalary() int {
	return es.fixedPay + es.joiningBonus + es.variablePay
}

type EmployeeSalaried struct {
	name string
	employeeSalary EmployeeSalary
}

func main() {
	var dushyant Employee
	fmt.Println(dushyant)

	dushyant.name = "Dushyant"
	dushyant.salary = 7500000
	dushyant.fullTime = true

	fmt.Println(dushyant)

	sapra := Employee{name: "Sapra", salary: 7500000, fullTime: true}
	fmt.Println(sapra)

	//Anonymous struct
	emp := struct {
		name   string
		salary int
	}{name: "employee", salary: 7500000}

	fmt.Println(emp)

	//Pointer to a struct
	ross := &Employee{
		name: "Pointer", salary: 7500000, fullTime: false,
	}

	fmt.Println(ross)
	fmt.Println(*ross)
	//We can access the fields of a struct pointer without dereferencing it first.
	//Go will take care of dereferencing a pointer under the hood.
	fmt.Println(ross.name)

	//Anonymous fields
	data := Data{"Data", 1, false}
	fmt.Println(data)

	//Nested Structs
	employeeData := EmployeeData{name: "Dushyant",
		salary: Salary{fixedPay: 6800000, variablePay: 700000, joiningBonus: 800000}, fullTime: true}

	fmt.Println(employeeData)
	fmt.Println("Fixed Pay is ", employeeData.salary.fixedPay)

	/*
		Promoted fields

		Note: If a nested anonymous struct has a same field (field name) that conflicts with the field
		name defined in the parent struct, then that field won’t get promoted.
		Only the non-conflicting fields will get promoted.
	*/
	employeeInfo := EmployeeInfo{name: "Employee",
		Salary: Salary{6800000, 800000, 700000}, fullTime: false}
	fmt.Println(employeeInfo)
	// Here all fields of Salary struct is directly available under EmployeeInfo struct.
	fmt.Println("JoiningBonus is ", employeeInfo.joiningBonus)

	es := EmployeeSalaried{"Dushyant", EmployeeSalary{6800000, 800000, 700000}}
	fmt.Println(es.employeeSalary.getSalary())
}
