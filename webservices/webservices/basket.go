package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkout", checkoutHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type OrderLine struct {
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}

// START_HANDLER OMIT
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	lines := []OrderLine{}
	dec := json.NewDecoder(r.Body) // HL
	err := dec.Decode(&lines)      // HL
	if err != nil {
		panic(err)
	}
	total := 0.0
	for _, line := range lines {
		total += line.Amount
	}

	ret := map[string]float64{"total": total}
	err = writeJSON(w, 200, ret)
	if err != nil {
		panic(err)
	}
}

// END_HANDLER OMIT

func writeJSON(w http.ResponseWriter, status int, d interface{}) error {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	if status > 0 {
		w.WriteHeader(status)
	}
	return enc.Encode(d)
}
