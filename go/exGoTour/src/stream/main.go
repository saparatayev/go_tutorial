package main
import "fmt"
 
func main(){
   intCh := make(chan int) 
    
   go factorial(7, intCh)
  
    for num := range intCh{
        fmt.Println(num)
    }
}
 
func factorial(n int, ch chan int){
    defer close(ch)
    result := 1
    for i := 1; i <= n; i++{
        result *= i
        ch <- result
    }
}