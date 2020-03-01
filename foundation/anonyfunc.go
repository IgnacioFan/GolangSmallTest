package main
import(
	"fmt"
)

func addOne() func() int {
	var x int
	
	// anonymous function will return integer
	return func() int {
		x++
		return x
	}
}

func main () {
	myFunc := addOne()
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())
}