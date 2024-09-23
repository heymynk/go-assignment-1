package main

import "fmt"

// Define an interface for polymorphism
type StaffMember interface {
	Pay() string
}

// Base struct for common staff members
type Staff struct {
	Name string
}

// Constructor for Staff
func NewStaff(name string) *Staff {
	return &Staff{Name: name}
}

// Method that all staff members implement
func (s *Staff) Pay() string {
	return s.Name + " is a staff member and is being paid."
}

// Employee type embedding Staff
type Employee struct {
	*Staff
}

// Constructor for Employee
func NewEmployee(name string) *Employee {
	return &Employee{NewStaff(name)}
}

// Overriding Pay method for Employee
func (e *Employee) Pay() string {
	return e.Name + " is an employee and is being paid."
}

// Manager type embedding Employee
type Manager struct {
	*Employee
}

// Constructor for Manager
func NewManager(name string) *Manager {
	return &Manager{NewEmployee(name)}
}

// Overriding Pay method for Manager
func (m *Manager) Pay() string {
	return m.Name + " is a manager and is being paid with a bonus."
}

func main() {
	// Array of StaffMember interface to store polymorphic types
	staffMembers := []StaffMember{
		NewEmployee("Alice"),
		NewManager("Bob"),
		NewEmployee("Charlie"),
	}

	// Iterate over staff members and call Pay method polymorphically
	for _, staff := range staffMembers {
		fmt.Println(staff.Pay())
	}
}
