package main
import "fmt"
 
type Vehicle interface{
    move()
}
 
func drive(vehicle Vehicle){
    vehicle.move()
}
 
type Car struct{ }
type Aircraft struct{}
type Ship struct{}
 
 
func (c Car) move(){
    fmt.Println("Автомобиль едет")
}
func (a Aircraft) move(){
    fmt.Println("Самолет летит")
}
func (s Ship) move(){
    fmt.Println("Корабль плывет")
}
 
func main() {
     
    tesla := Car{}
	boing := Aircraft{}
    drive(tesla)
	drive(boing)
	drive(Ship{})
}