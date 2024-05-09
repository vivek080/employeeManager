package service

import (
	"employeeManager/gosrc/model"
	"testing"
)

// TestNewEmployeeStore tests the initialization of a new EmployeeStore
func TestNewEmployeeStore(t *testing.T) {
	store := NewEmployeeStore()
	if store == nil {
		t.Error("NewEmployeeStore() returned nil")
	}
	if len(store.Employees) != 0 {
		t.Errorf("Expected empty store, got %v", store.Employees)
	}
}

// TestCreateEmployee tests adding a new employee
func TestCreateEmployee(t *testing.T) {
	store := NewEmployeeStore()
	emp := model.Employee{Name: "ABC", Position: "Software Developer", Salary: 40000}
	id := store.CreateEmployee(emp)

	if id != 1 {
		t.Errorf("Expected first employee ID to be 1, got %d", id)
	}

	if len(store.Employees) != 1 {
		t.Errorf("Expected store to have 1 employee, got %d", len(store.Employees))
	}

	if _, exists := store.Employees[id]; !exists {
		t.Errorf("Employee with ID %d was not found in store", id)
	}
}

// TestGetEmployeeByID tests retrieving an employee by ID
func TestGetEmployeeByID(t *testing.T) {
	store := NewEmployeeStore()
	emp := model.Employee{Name: "Vicky", Position: "Manager", Salary: 70000}
	id := store.CreateEmployee(emp)

	retrievedEmp, exists := store.GetEmployeeByID(id)
	if !exists {
		t.Errorf("Expected to find employee with ID %d", id)
	}
	if retrievedEmp.Name != emp.Name {
		t.Errorf("Expected Name %s, got %s", emp.Name, retrievedEmp.Name)
	}
}

// TestUpdateEmployee tests updating an employee's details
func TestUpdateEmployee(t *testing.T) {
	store := NewEmployeeStore()
	emp := model.Employee{Name: "Sonic", Position: "CEO", Salary: 90000}
	id := store.CreateEmployee(emp)

	// Update employee
	emp.ID = id
	emp.Name = "Sam"
	success := store.UpdateEmployee(emp)
	if !success {
		t.Errorf("Failed to update employee with ID %d", id)
	}

	updatedEmp, exists := store.GetEmployeeByID(id)
	if !exists || updatedEmp.Name != "Sam" {
		t.Errorf("Update failed, expected Name %s, got %s", "Sam", updatedEmp.Name)
	}
}

// TestDeleteEmployee tests the deletion of an employee
func TestDeleteEmployee(t *testing.T) {
	store := NewEmployeeStore()
	emp := model.Employee{Name: "Apex", Position: "HR Manager", Salary: 50000}
	id := store.CreateEmployee(emp)

	success := store.DeleteEmployee(id)
	if !success {
		t.Errorf("Failed to delete employee with ID %d", id)
	}

	_, exists := store.GetEmployeeByID(id)
	if exists {
		t.Error("Employee was found after deletion")
	}
}

// TestListEmployees tests pagination functionality
func TestListEmployees(t *testing.T) {
	store := NewEmployeeStore()
	for i := 0; i < 10; i++ {
		emp := model.Employee{Name: "Employee", Position: "Position", Salary: 50000}
		store.CreateEmployee(emp)
	}

	employees := store.ListEmployees(1, 5)
	if len(employees) != 5 {
		t.Errorf("Expected 5 employees per page, got %d", len(employees))
	}
}
