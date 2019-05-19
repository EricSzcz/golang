package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Nova")
	fmt.Print(" linha.")

	x := 3.141516
	//fmt.Println("O valorde x é:" + x) nao funciona

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é", xs)
	fmt.Println("O valor de x é", x, "!!!")

	fmt.Printf("O valor de x é %.2f", x) // formatação com duas casa decimais

	a := 1
	b := 1.999
	c := false
	d := "Opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // impressão dos tipos
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
