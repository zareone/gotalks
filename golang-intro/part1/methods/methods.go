package main

import (
	"fmt"
	"strings"
)

// START OMIT
type email string

func (e email) validate() error {
	if !strings.Contains(string(e), "@") {
		return fmt.Errorf("email validate: %s is missing @", e)
	}

	return nil
}

func main() {
	fmt.Println(
		email("camilo@gmail.com").validate(),
		email("josepogmail.com").validate(),
	)
}

// END OMIT
