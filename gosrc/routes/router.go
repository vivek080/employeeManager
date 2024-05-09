package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"employeeManager/gosrc/model"
	"employeeManager/gosrc/service"

	"github.com/gorilla/mux"
)

func RouteCalls(router *mux.Router, store *service.EmployeeStore) {
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			pageStr, pageSizeStr := r.URL.Query().Get("page"), r.URL.Query().Get("pagesize")
			page, err1 := strconv.Atoi(pageStr)
			pageSize, err2 := strconv.Atoi(pageSizeStr)
			if err1 != nil || err2 != nil || page < 1 || pageSize < 1 {
				http.Error(w, `{"error":"Invalid pagination parameters"}`, http.StatusBadRequest)
				return
			}
			employees := store.ListEmployees(page, pageSize)
			json.NewEncoder(w).Encode(employees)
		case http.MethodPost:
			var emp model.Employee
			if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id := store.CreateEmployee(emp)
			json.NewEncoder(w).Encode(map[string]int{"id": id})
		}
	}).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, `{"error":"Invalid employee ID"}`, http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			if emp, exists := store.GetEmployeeByID(id); exists {
				json.NewEncoder(w).Encode(emp)
			} else {
				http.Error(w, `{"error":"Employee ID not found"}`, http.StatusNotFound)
			}
		case http.MethodPut:
			var emp model.Employee
			if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			emp.ID = id
			if !store.UpdateEmployee(emp) {
				http.Error(w, `{"error":"Employee ID not found"}`, http.StatusNotFound)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"message": "Employee Data updated successfully"})
			}
		case http.MethodDelete:
			if !store.DeleteEmployee(id) {
				http.Error(w, `{"error":"Employee ID not found"}`, http.StatusNotFound)
			} else {
				json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("Employee ID: %d deleted successfully", id)})
			}
		}
	}).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
}
