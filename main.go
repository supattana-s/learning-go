package main

import (
	"fmt"

	"github.com/supattana-s/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	fmt.Println(s3)

	s4 := puppy.BigBarks()
	fmt.Println(s4)
}	