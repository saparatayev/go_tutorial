package main
import (
    "fmt"
    "os"
)
 
// type person struct { 
//    name string
//    age int32
//    weight float64
// }
func main() {
    filename := "hello2.txt"
    writeData(filename)
    readData(filename)
}
 
func writeData(filename string){
    // начальные данные
    var name string = "Tom"
    var age int = 24
     
    file, err := os.Create(filename)
    if err != nil { 
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close() 
     
    fmt.Fprintln(file, name)
    fmt.Fprintln(file, age)
}
func readData(filename string){
     
    var name string
    var age int
     
    file, err := os.Open(filename)
    if err != nil{
        fmt.Println(err) 
        os.Exit(1)
    }
    defer file.Close()
     
    fmt.Fscanln(file, &name) 
    fmt.Fscanln(file, &age) 
    fmt.Println(name, age)
}