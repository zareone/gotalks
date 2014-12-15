package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START_RATE OMIT
type rate uint8 // Domain values are 0-5

func (r rate) MarshalJSON() ([]byte, error) {
	ratings := []string{
		`"awful"`,
		`"bad"`,
		`"meh"`,
		`"good"`,
		`"great"`,
		`"awesome"`,
	}
	for min, desc := range ratings {
		if int(r) <= min {
			return []byte(desc), nil
		}
	}
	return []byte(`"awful"`), nil
}

// END_RATE OMIT

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", langsHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// START_LANGS OMIT
var langs = []struct {
	Name            string `json:"name"`
	Year            int    `json:"appearedIn"`
	Compiled        bool   `json:"-"` // HL
	WebFriendliness rate   `json:"webFriendliness"`
	Performance     rate   `json:"performance"`
}{
	{"C", 1972, true, 1, 5},
	{"C++", 1983, true, 2, 5},
	{"Python", 1991, false, 4, 3},
	{"Go", 2009, true, 5, 4},
	{"Rust", 2012, true, 4, 5},
}

// END_LANGS OMIT

// START_HANDLER OMIT
func langsHandler(w http.ResponseWriter, r *http.Request) {
	seq, err := json.Marshal(langs) // HL
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json") // HL
	w.Write(seq)
}

// END_HANDLER OMIT
