package main

import (
	"fmt"

	"github.com/EricSzcz/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	//fmt.Println(area._TrianguloEq(6.0)) Não sera encontrado pois é privado
}
