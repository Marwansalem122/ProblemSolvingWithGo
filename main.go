package main

import "math"

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	println(mySqrt(2))
	println(mySqrt(8))
}

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}
