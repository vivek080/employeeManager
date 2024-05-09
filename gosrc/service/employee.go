// /gosrc/service/employee.go
package service

import (
	"employeeManager/gosrc/model"
	"sync"
)

type EmployeeStore struct {
	sync.RWMutex
	Employees map[int]model.Employee
	UniqueID  int
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		Employees: make(map[int]model.Employee),
		UniqueID:  1,
	}
}

// CreateEmployee creates a new employee
func (store *EmployeeStore) CreateEmployee(emp model.Employee) int {
	store.Lock()
	defer store.Unlock()

	emp.ID = store.UniqueID
	store.Employees[emp.ID] = emp
	store.UniqueID++

	return emp.ID
}

// GetEmployeeByID gets the details of the employee using the given ID
func (store *EmployeeStore) GetEmployeeByID(id int) (model.Employee, bool) {
	store.RLock()
	defer store.RUnlock()

	emp, exists := store.Employees[id]

	return emp, exists
}

// UpdateEmployee updates the data of the employee with the given ID
func (store *EmployeeStore) UpdateEmployee(emp model.Employee) bool {
	store.Lock()
	defer store.Unlock()

	_, exists := store.Employees[emp.ID]
	if exists {
		store.Employees[emp.ID] = emp
	}

	return exists
}

// DeleteEmployee deletes the employee with the given ID
func (store *EmployeeStore) DeleteEmployee(id int) bool {
	store.Lock()
	defer store.Unlock()

	_, exists := store.Employees[id]
	if exists {
		delete(store.Employees, id)
	}

	return exists
}

// ListEmployees list the employee list with pagination
func (store *EmployeeStore) ListEmployees(page, pageSize int) []model.Employee {
	store.RLock()
	defer store.RUnlock()

	var employees []model.Employee
	for _, emp := range store.Employees {
		employees = append(employees, emp)
	}
	start := (page - 1) * pageSize
	if start > len(employees) {
		return nil
	}
	end := start + pageSize
	if end > len(employees) {
		end = len(employees)
	}

	return employees[start:end]
}
