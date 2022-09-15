package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":8888"

type EmployeeModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

var employees = []EmployeeModel{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, e *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": http.StatusText(http.StatusMethodNotAllowed),
			})
		}
		var req EmployeeModel
		var err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": http.StatusText(http.StatusBadRequest),
			})
		}

		req.Id = len(employees) + 1
		employees = append(employees, req)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  http.StatusText(http.StatusOK),
			"employee": employees,
		})
		return
	})
	log.Println("RUN IN ", port)
	http.ListenAndServe(port, nil)
}
