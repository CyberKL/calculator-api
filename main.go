package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/subtract", handleSubtract)
	http.HandleFunc("/multiply", handleMultiply)
	http.HandleFunc("/divide", handleDivide)
	http.HandleFunc("/sum", handleSum)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

type Response struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

type Operands struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

func getOperands(r *http.Request) (int, int, error) {
	operands := &Operands{}
	if err := json.NewDecoder(r.Body).Decode(operands); err != nil {
		return 0, 0, err
	}

	// Validate that the numbers are whole numbers
	if operands.Number1 != float64(int(operands.Number1)) || operands.Number2 != float64(int(operands.Number2)) {
		return 0, 0, errors.New("Operands must be whole numbers")
	}

	return int(operands.Number1), int(operands.Number2), nil
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

	number1, number2, err := getOperands(r)
	if err != nil {
		if err.Error() == "Operands must be whole numbers" {
			writeJSON(w, http.StatusBadRequest, Response{Error: "Operands must be whole numbers"})
		} else {
			log.Println(err)
			writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		}
		return
	}

	result := Response{Result: number1 + number2}
	writeJSON(w, http.StatusOK, result)
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	number1, number2, err := getOperands(r)
	if err != nil {
		if err.Error() == "Operands must be whole numbers" {
			writeJSON(w, http.StatusBadRequest, Response{Error: "Operands must be whole numbers"})
		} else {
			log.Println(err)
			writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		}
		return
	}

	result := Response{Result: number1 - number2}
	writeJSON(w, http.StatusOK, result)
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	number1, number2, err := getOperands(r)
	if err != nil {
		if err.Error() == "Operands must be whole numbers" {
			writeJSON(w, http.StatusBadRequest, Response{Error: "Operands must be whole numbers"})
		} else {
			log.Println(err)
			writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		}
		return
	}

	result := Response{Result: number1 * number2}
	writeJSON(w, http.StatusOK, result)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	number1, number2, err := getOperands(r)
	if err != nil {
		if err.Error() == "Operands must be whole numbers" {
			writeJSON(w, http.StatusBadRequest, Response{Error: "Operands must be whole numbers"})
		} else {
			log.Println(err)
			writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		}
		return
	}

	if number2 == 0 {
		writeJSON(w, http.StatusBadRequest, Response{Error: "Division by zero is not allowed"})
		return
	}

	result := Response{Result: number1 / number2}
	writeJSON(w, http.StatusOK, result)
}

func getNumbers(r *http.Request) ([]int, error) {
	var numbers []int
	if err := json.NewDecoder(r.Body).Decode(&numbers); err != nil {
		return nil, err
	}

	return numbers, nil
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "Method not allowed"})
		return
	}

	numbers, err := getNumbers(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, Response{Error: "Invalid JSON payload"})
		return
	}

	result := 0
	for _, number := range numbers {
		result += number
	}
	
	writeJSON(w, http.StatusOK, Response{Result: result})
}