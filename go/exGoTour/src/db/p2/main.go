package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
type product struct{
    id int
    model string
    company string
    price int
}
func main() { 
    db, err := sql.Open("mysql", "root:@/productdb")
     
    if err != nil {
        panic(err)
    } 
    defer db.Close()
    rows, err := db.Query("select model, price from productdb.Products")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    products := []product{}
     
    for rows.Next(){
        p := product{}
        err := rows.Scan(&p.model, &p.price)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
    for _, p := range products{
        fmt.Println(p.id, p.model, p.company, p.price)
    }
}