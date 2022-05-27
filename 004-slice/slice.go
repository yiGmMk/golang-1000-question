package slice

import (
	"log"
	"math/rand"
)

// we have 1 in slice,and range the list,every time we add a pseudo-random number to the list
// what is value of n?
// and len(list) = ? when the end
func AppendIteminRange() {
	list := []int{1}
	add := 0
	for n := len(list); n > 0 && n < 10; {
		add++
		list = append(list, rand.Intn(100))
		log.Println("n=", n, " add=", add, " list:", list)
		if add > 20 {
			break
		}
	}
}
