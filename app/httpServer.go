package main

import (
	"encoding/json"
	"net/http"
)

func pdfCreationHandler(w http.ResponseWriter, r *http.Request) {

	var req = request{}
	switch contentType := r.Header.Get("Content-type"); contentType {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			//TODO log error here
			return
		}
	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
		//TODO log error here
		return
	}

}
