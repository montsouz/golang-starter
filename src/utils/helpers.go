package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ResponseError to respond some errors
func ResponseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

//ResponseJSON to respond with some JSON
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

//ReadBodyAndUnmarshalData helper function
func ReadBodyAndUnmarshalData(r *http.Request, v interface{}) error {
	// Read body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// Read post
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
