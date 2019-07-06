package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	a , _ :=json.Marshal( map[string]int{"213":2,"2133":3} )


	fmt.Println(a)
}

