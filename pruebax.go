package main

import (
	"fmt"
	"sync"
	"time"
)

type palillo struct{ sync.Mutex }

type filosofo struct {
	id int
	pIzq, pDer *palillo
}
func (p filosofo)comer() {
	for i:=0;i<3;i++{
		p.pIzq.Lock()
		p.pDer.Lock()
		fmt.Println("Filosofo #",p.id," esta comenzando  a comer")
		time.Sleep(time.Second)
		fmt.Println("Filosofo #",p.id," esta terminando de comer")
		p.pIzq.Unlock()
		p.pDer.Unlock()
		time.Sleep(time.Second)
	}
	cond.Done()
}


const N = 5
var cond sync.WaitGroup
func main() {


	// slice palillos
	palillos := make([]*palillo, N)
	filosofos := make([]*filosofo, N)
	for i:=0; i<N; i++{
		palillos[i] = new(palillo)
	}
	for i:=0; i<N; i++{
		filosofos[i] = &filosofo{id: i, pIzq: palillos[i], pDer: palillos[(i+1)%N]}
	}

	for i:=0; i<N; i++{
		cond.Add(1)
		go filosofos[i].comer()
	}
	cond.Wait()

}