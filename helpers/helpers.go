package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("Error Creating JSON")
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error Writing to response")
		return err
	}
	return nil

}
