package main

import (
	"errors"
	"fmt"
)

func division(a float32, b float32) (float32, error) {
	var err error = nil
	if b == 0 {
		err = errors.New("Division indefinida")
		return 0, err
	}
	resultado := a / b

	return resultado, err
}

func main() {
	div, err := division(5, 0)
	if err != nil {
		fmt.Println("No se puede dividir. Error:", err.Error())

	}

	fmt.Println("division: ", div)

}
