package main

import (
	"bootcamp/bigint/bigint"
	"fmt"
)

func main() {

	a, err := bigint.NewInt("1")
	if err != nil {
		panic(err)
	}
	b, err := bigint.NewInt("11231231231231231231231231238")
	if err != nil {
		panic(err)
	}

	err = a.Set("2")
	if err != nil {
		panic(err)
	}

	// c := bigint.Add(a, b)
	// d := bigint.Sub(a, b)
	e := bigint.Multiply(a, b)
	// f := bigint.Mod(a, b)
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	fmt.Println(e)
	// fmt.Println(f)
}
