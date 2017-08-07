package main

import "fmt"

type metodos interface {
	imprime()
}
type monitor struct {
	polegada int
}
type computador struct {
	monitor
	marca       string
	ram         int
	processador string
}

func imprime(c computador) {
	fmt.Printf("Marca: %s Ram: %d, Processador: %s Polegadas %v", c.marca, c.ram, c.processador, c.polegada)

}
func main() {
	c := new(computador)
	c.marca = "Dell"
	c.ram = 4
	c.processador = "i3"
	c.polegada = 22

	imprime(*c)
}
