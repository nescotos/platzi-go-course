package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	operador := "p"
	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operador1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(operador1)
	}
	operador2, _ := strconv.Atoi(valores[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	default:
		fmt.Println(operador, "No esta soportado")
	}

}
