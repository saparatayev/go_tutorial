// package main
// import (
//     "fmt"
// )

// type phoneWriter struct{ }
 
// func (p phoneWriter) Write(bs []byte) (int, error){
//     if len(bs) == 0 { 
//          return 0, nil 
//    }
//    for i := 0; i < len(bs); i++{
//         if(bs[i] >= '0' && bs[i] <= '9'){
//             fmt.Print(string(bs[i]))
//         }
//     }
//     fmt.Println()
//     return len(bs), nil
// }
 
// func main() {
//     writer := phoneWriter{}
	
//     fmt.Fprint(writer, "+1(234)567 9010")
//     fmt.Fprintln(writer, "+2-345-678-12-35")
// }

package main
import (
    "fmt"
    "os"
)
 
func main() {
    file, err := os.Create("confeve.txt")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
    fmt.Fprint(file, "Сегодня ")
    fmt.Fprintln(file, "hot погода")
}