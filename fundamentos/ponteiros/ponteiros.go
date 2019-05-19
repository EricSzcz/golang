package main

import "fmt"

func main() {
	i := 1

	var p *int

	p = &i // pegando o endereço da variavel
	*p++   // desrefenciando para pegar o valor ao inves da referencia de memoria
	i++    // mesma coisa da linha acima

	//Go não tem aritimetica de ponteiros
	//p++

	fmt.Println(p, i, *p, &i)
}
