package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chapter2/ohm/power"
	"chapter2/ohm/resist"
	"chapter2/ohm/volt"
)

type telemetry struct {
	I float64   `json:"i"`
	R []float64 `json:"r"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data telemetry
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error decoding data: %v", err)
			return
		}
		p := power.Piv(data.I, volt.Vir(data.I, resist.Rser(data.R...)))
		fmt.Fprintf(w, "{power: %f}", p)
	})
	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
