package main

import "fmt"
import "os"
import "errors"
import "strconv"
import "math"
var ErrDividedO error = errors.New("division by zero is impossible") 
	var NotNumber error = errors.New("Not A numericale number.")
		
	func add(nb,nb1 float64) float64 {
	return nb + nb1
}

func sub(nb,nb1 float64) float64 {
	return nb - nb1
}


func mul(nb,nb1 float64) float64 {
	return nb * nb1
}


func div(nb,nb1 float64) (float64, error) {
	if nb1 == 0 {
		return 0, ErrDividedO
	}
	return nb / nb1, nil
}
func mod(nb,nb1 float64) (float64, error) {
	if nb1 == 0 {
		return 0, ErrDividedO
	}
	return math.Mod(nb,nb1), nil
}


func main() {
	if len(os.Args)-1 != 3 {
		fmt.Println("only three arguments allowed")
				os.Exit(1)
	}
		jmi, opp ,nbr := os.Args[1], os.Args[2] ,os.Args[3]
			jm, err := strconv.ParseFloat(jmi,64)
			if err != nil {
				fmt.Println(NotNumber)
					os.Exit(1)
			}
			nb, err	:= strconv.ParseFloat(nbr,64)
				if err != nil {
					fmt.Println(NotNumber)
						os.Exit(1)
				}
			switch opp {
			case "+": 
			fmt.Printf("%.2f", add(jm,nb))
			case "-":
			fmt.Printf("%.2f", sub(jm,nb))
			case "*": 
			fmt.Printf("%.2f", mul(jm, nb))
			case "/": 
			res, err := div(jm , nb)
			if err != nil {
				fmt.Println(ErrDividedO)
					os.Exit(1)
				} else {
				fmt.Printf("%.2f", res)
			}
			case "%": 
			resu, err := mod(jm, nb)
			if err != nil {
				fmt.Println(ErrDividedO)
					os.Exit(1)
			}
			fmt.Printf("%.2f", resu)
			default:
				fmt.Println("Function is not understandble. Pls use the allowed opperateurs:\n+ | - | * | / | %.")
			}		
				

		
}