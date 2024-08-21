package main

import "fmt"

var m = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func main() {
	fmt.Println(romanToInt("XV"))
}

// percorre os números de trás para frente
// se o número atual for menor que o anterior, ele decrementa o valor final
// se for maior, ele incrementa
func romanToInt(s string) int {
	var v, cv, lv int
	fmt.Println(v)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(v)
		cv = m[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
		fmt.Println(v)
	}
	return v
}
