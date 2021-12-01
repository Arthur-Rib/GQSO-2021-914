package main

import (
	"fmt"
	"strconv"
)

func soma(x, y string) (float64, error) {
	op1, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, fmt.Errorf("a string %s nao pode ser transformada em inteiro", x)
	}
	op2, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, fmt.Errorf("a string %s nao pode ser transformada em inteiro", y)
	}
	return op1 + op2, nil
}
