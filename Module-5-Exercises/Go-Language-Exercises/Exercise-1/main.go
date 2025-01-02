package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID   int
	Name string
	Age  int
	Dept string
}

var employees []Employee

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Add employee
func addEmployee(id int, name string, age int, dept string) error {
	for _, emp := range employees {
		if emp.ID == id {
			return ValidationError{"ID", "Must be unique!"}
		}
	}

	if age <= 18 {
		return ValidationError{"Age", "Must be greater than 18"}
	}

	employees = append(employees, Employee{ID: id, Name: name, Age: age, Dept: dept})
	return nil
}

// Search employee by ID or name
func searchEmployee(query string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprint(emp.ID) == query || emp.Name == query {
			return &emp, nil
		}
	}
	return nil, errors.New("Employee not found!")
}

// List employees by department
func listEmployeesByDepartment(dept string) {
	fmt.Printf("Employees in %s department:\n", dept)
	for _, emp := range employees {
		if emp.Dept == dept {
			fmt.Printf("ID: %d, Name: %s, Age: %d\n", emp.ID, emp.Name, emp.Age)
		}
	}
}

// Count employees in a department
func countEmployees(dept string) int {
	count := 0
	for _, emp := range employees {
		if emp.Dept == dept {
			count++
		}
	}
	return count
}

func main() {
	const hrDept = "HR"
	const itDept = "IT"
	const finDept = "Finance"

	// Adding employees
	if err := addEmployee(1001, "Arjun Singh", 25, hrDept); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	if err := addEmployee(1002, "Priya Kumari", 30, itDept); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	if err := addEmployee(1003, "Chetna Pande", 42, finDept); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	if err := addEmployee(1004, "Vikash Sharma", 34, finDept); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	if err := addEmployee(1005, "Mehali Samanta", 22, itDept); err != nil {
		fmt.Println("Error adding employee:", err)
	}
	// Searching employees
	emp, err := searchEmployee("Priya Kumari")
	if err != nil {
		fmt.Println("Employee does not Exist", err)
	} else {
		fmt.Printf("Found employee: %+v\n", *emp)
	}

	// Listing employees by department
	listEmployeesByDepartment(itDept)

	// Counting employees
	fmt.Printf("Number of employees in %s: %d\n", finDept, countEmployees(finDept))
}
