package main
import(
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

)
func main() {
	fmt.Printf("\n%s\n",isAnagram(os.Args))
}

func isAnagram(args []string) string {
	if len(args) != 3 {
		return "bad parameters"
	}
	a1 := strings.Split(args[1], "")
	a2 := strings.Split(args[2], "")
	
	sort.Strings(a1)
	sort.Strings(a2)
	if reflect.DeepEqual(a1, a2) {
		return "YES"
	}
	return "NO"
}