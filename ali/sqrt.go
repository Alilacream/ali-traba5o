package ali

import (
	"fmt"
	"os"
)

/*Baby lEVEL*/
/*func Aqrt(nb int) int {
	var jmi int = 0
	for jmi * jmi < nb {
		jmi++
	}
	return jmi
}*/
/*Advanced Level : Newton-raphson*/
func Sqrt(nb int64) int64 {
	if nb < 0 {
		fmt.Println("The number must be positive")
		os.Exit(1)
	}
	x := nb / 2
	var l float64 = 0.00001
	y := x + nb/x
	for {
		guess := (x + y) / 2
		if float64(abs(guess-x)) < l {
			break
		}
		x = guess
	}
	return x
}
func abs(nb int64) int64 {
	if nb < 0 {
		return -nb
	}
	return nb
}
