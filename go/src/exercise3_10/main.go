package main
import(
	"fmt"
	"os"
	"strings"
	// "strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var s2 string
	var dot int
	if dot = strings.LastIndex(s, "."); dot >= 0 {
		s2 = s[dot:]
	} else {
		dot = len(s)
	}
	for i,j := dot-1,0; i >=0; i-- {

		if string(s[i]) == "-" {
			s2 = string(s[i]) + s2
			break
		}

		if j%3 == 0 && j != 0 {
			s2 = "," + s2
		}

		s2 = string(s[i]) + s2
		
		j++
	}
	return s2
}