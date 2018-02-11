package main

import (
	"fmt"

	"github.com/auriou/gosample/libs/calculatrice"
)

func main() {
	x := calculatrice.Add(4, 5)
	fmt.Println("resultat", x)
}
