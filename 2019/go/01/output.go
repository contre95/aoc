package main
import (
	"fmt"
)

func calc_fuel(mass int) (fuel int){
	fuel = mass / 3 - 2
	return fuel
}

func main() {
	fmt.Println(calc_fuel(1969))
	fmt.Println(10/6)
}
