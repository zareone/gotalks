package main

// START_HANDLER OMIT
func langsHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w) // HL
	w.Header().Set("Content-Type", "application/json")

	err := enc.Encode(langs) // HL
	if err != nil {
		panic(err)
	}
}

// END_HANDLER OMIT
