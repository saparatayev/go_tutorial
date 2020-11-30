package main
import (
"fmt"
"io"
)
 
type phoneReader string
 
func (ph phoneReader) Read(p []byte) (int, error){
    count := 0
    for i := 0; i < len(ph); i++{
        if(ph[i] >= '0' && ph[i] <= '9'){
            p[count] = ph[i]
            count++
        }
    }
    return count, io.EOF
}
 
func main() { 
    var phone1 phoneReader = "+1(234)567 9010"
    phone2 := phoneReader("+2-345-678-12-35")
     
    buffer := make([]byte, len(phone1))
    count, _ := phone1.Read(buffer)
    fmt.Println(string(buffer),"read", count)     // 12345679010
     
    buffer = make([]byte, len(phone2))
    count, _ = phone2.Read(buffer)
    fmt.Println(string(buffer),"read", count)     // 23456781235
}

// package main
// import "fmt"
 
// type mile int
 
// func main() {
     
//     distance := mile(5)
//     fmt.Println(distance)
//     distance += 5
//     fmt.Println(distance)
// }