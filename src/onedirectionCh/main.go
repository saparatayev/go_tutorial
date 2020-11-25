package main
import "fmt"
 
func main() {
     
    intCh := make(chan int, 2) 
    go factorial(5, intCh)
    fmt.Println(<-intCh)
    fmt.Println("The End")
}
 
func factorial(n int, ch chan<- int){
     
    result := 1
    for i := 1; i <= n; i++{
        result *= i
    }
    ch <- result
}