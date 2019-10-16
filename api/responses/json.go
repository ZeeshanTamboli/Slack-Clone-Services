package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON : Return data by converting it to JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR : Returns error. If err is given then returns "error" object else returns Bad Request
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	type CustomError struct {
		Error string `json:"error"`
	}
	if err != nil {
		JSON(w, statusCode, CustomError{
			Error: err.Error(),
		})
	}

	JSON(w, http.StatusBadRequest, nil)
}
