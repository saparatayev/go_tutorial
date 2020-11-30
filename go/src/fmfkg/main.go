package main

import (
	"fmt"
	"os"
	"strconv"

	"./convpack"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := convpack.Meters(v)
		f := convpack.Feets(v)
		p := convpack.Pounds(v)
		k := convpack.Kilograms(v)
		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n\n",
			m, convpack.MToF(m), f, convpack.FToM(f),
			k, convpack.KToP(k), p, convpack.PToK(p))
	}
}