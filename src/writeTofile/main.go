// package main
// import (
//     "fmt"
//     "os"
// )
 
// func main() {
//     text := "Hello Gold!"
//     file, err := os.Create("hello.txt")
     
//     if err != nil{
//         fmt.Println("Unable to create file:", err) 
//         os.Exit(1) 
//     }
//     defer file.Close() 
//     file.WriteString(text)
     
//     fmt.Println("Done.")
// }

package main
import (
    "fmt"
    "os"
)
 
func main() {
    data := []byte("Hello Bold!")
    file, err := os.Create("hello.bin")
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
    defer file.Close() 
    file.Write(data)
     
    fmt.Println("Done.")
}