package main

import (
	"fmt"
)

const M = 3

func main() {
	var pilha [M]int
	topo := 0

	insere(&pilha, &topo, 4)
	fmt.Println(pilha)

	insere(&pilha, &topo, 2)
	fmt.Println(pilha)

	insere(&pilha, &topo, 12)
	fmt.Println(pilha)

	insere(&pilha, &topo, 1)
	fmt.Println(pilha)

	remove(&pilha, &topo)
	fmt.Println(pilha)

}

func insere(p *[M]int, topo *int, valor int) {
	if *topo == M {
		fmt.Println("Overflow")
		return
	}

	p[*topo] = valor
	*topo++

}

func remove(p *[M]int, topo *int) int {
	if *topo == 0 {
		fmt.Println("Underflow")
		return -1
	}

	*topo--
	valorRemovido := p[*topo]
	p[*topo] = 0
	return valorRemovido
}
