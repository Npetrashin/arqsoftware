package main

import (
	"errors"
	"fmt"
)

func Division(a int, b int) (int, error) {
	var err error = nil
	if b == 0 {
		err = errors.New("Division indefinida")
		return 0, err
	}

	return a/b, nil
}
func main() {
	div, err := Division(12, 4)
	if err != nil {
		fmt.Println("No se puede dividir. Error:", err.Error())

	}

	fmt.Println("division: ", div)
}