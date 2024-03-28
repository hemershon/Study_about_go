package main

import (
	"fmt"
	"strconv"
)

func Hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func sum(a, b int) int {
	return a + b
}
func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {

		return
	}

	total = a + i

	return
}

func main() {
	hello("Hemershon Silva")
	fmt.Println("total", sum(10, 5))
}
