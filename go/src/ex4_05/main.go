package main
import (
	"./unique"
	"os"
	"fmt"
)

func main() {
	fmt.Println(unique.Unique(os.Args[1:]))
}