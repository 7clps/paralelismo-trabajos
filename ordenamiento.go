package main

import (
	"fmt"
	"sync"
)
var cond sync.WaitGroup

func main() {
	x := []int{19, 3, 2, 5, 80, 88, 77, 99, 50, 7, 15, 16, 1, 35, 28, 48, 56, 54, 31, 37, 11, 14}
	cond.Add(5)
	fmt.Println("Matriz: ",x, len(x))
	ordenar(x)
	orden(x)
	cond.Wait()
}
func ordenar(x []int) {
	size := len(x)/4
	a,b,c:= size, 2*size, 3*size
	go orden(x[:a])
	go orden(x[a:b])
	go orden(x[b:c])
	go orden(x[c:])

}
func orden(x []int) {
	size := len(x)
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if x[j] < x[j-1] {
				x[j], x[j-1] = x[j-1], x[j]
			}
		}
	}
	fmt.Println(x)
	cond.Done()
}

