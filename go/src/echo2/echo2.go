package main
import (
	"fmt"
	"os"
	// "io"
)

// Create Stdout for testing
// var out io.Writer = os.Stdout

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)
	for index, arg := range os.Args {
		fmt.Printf("[%d] %s\n", index, arg)
	}
}