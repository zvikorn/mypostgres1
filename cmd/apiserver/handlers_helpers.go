package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// writeJSONResponse writes payload to response writer as json string
func writeJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var err error
	if code == http.StatusNoContent {
		return
	}

	var response []byte
	if payload != nil {
		response, err = json.Marshal(payload)
	} else {
		response, err = json.Marshal([]byte{})
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't marshal JSON response: %v", err)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		fmt.Println("error")
		fmt.Fprintf(os.Stderr, "Can't write JSON response: %v", err)
		return
	}
	_, err = w.Write([]byte("\n"))
}
