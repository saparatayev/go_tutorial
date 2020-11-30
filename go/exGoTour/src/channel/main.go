// package main
// import "fmt"
 
// func main() {
     
//     intCh := make(chan int) 
     
//     go func(){
//             fmt.Println("Go routine starts")
//             intCh <- 5 // блокировка, пока данные не будут получены функцией main
// 	}()
// 	// fmt.Scanln()
//     fmt.Println(<-intCh) // получение данных из канала
//     fmt.Println("The End")
// }

package main
import "fmt"
 
func main() {
     
    intCh := make(chan int) 
     
    go factorial(5, intCh)  // вызов горутины
    fmt.Println(<-intCh) // получение данных из канала
    fmt.Println("The End")
}
 
func factorial(n int, ch chan int){
     
    result := 1
    for i := 1; i <= n; i++{
        result *= i
    }
    fmt.Println(n, "-", result)
     
    ch <- result     // отправка данных в канал
}