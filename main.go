package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/subtract", handleSubtract)
	http.HandleFunc("/multiply", handleMultiply)
	http.HandleFunc("/divide", handleDivide)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

type Operands struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func getOperands(r *http.Request) (float64, float64, error) {
	operands := &Operands{}
	if err := json.NewDecoder(r.Body).Decode(operands); err != nil {
		return 0, 0, err
	}

	return operands.A, operands.B, nil
}

func writeJSON(w http.ResponseWriter, s int, v Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	a, b, err := getOperands(r)
	if err != nil {
		log.Println(err)
		writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		return
	}

	result := Response{Result: a + b}
	writeJSON(w, http.StatusOK, result)
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	a, b, err := getOperands(r)
	if err != nil {
		log.Println(err)
		writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		return
	}

	result := Response{Result: a - b}
	writeJSON(w, http.StatusOK, result)
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	a, b, err := getOperands(r)
	if err != nil {
		log.Println(err)
		writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		return
	}

	result := Response{Result: a * b}
	writeJSON(w, http.StatusOK, result)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	a, b, err := getOperands(r)
	if err != nil {
		log.Println(err)
		writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		return
	}

	if b == 0 {
		writeJSON(w, http.StatusBadRequest, Response{Error: "Division by zero is not allowed"})
		return
	}

	result := Response{Result: a / b}
	writeJSON(w, http.StatusOK, result)
}
