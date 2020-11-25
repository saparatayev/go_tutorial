package main
import (
    "fmt"
    "os"
    "io"
)
 
func main() {
    file, err := os.Open("../writeTofile/hello.txt")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
     
    io.Copy(os.Stdout, file)
}