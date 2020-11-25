package main
import (
    "fmt"
    "os"
    "bufio"
    "io"
)
func main(){
    file, err := os.Open("../../exGoTour.txt") 
    if err != nil { 
         fmt.Println("Unable to open file:", err) 
         return
    } 
    defer file.Close() 
  
    reader := bufio.NewReader(file) 
    for { 
        line, err := reader.ReadString('e') 
        if err != nil { 
            if err == io.EOF { 
                break
            } else { 
                fmt.Println(err) 
                return
            } 
        } 
        fmt.Print(line) 
    }
}