package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() { 
    db, err := sql.Open("mysql", "root:@/productdb")
     
    if err != nil {
        panic(err)
    } 
    defer db.Close()
     
    result, err := db.Exec("insert into productdb.Products (model, company, price) values ('Pixel 2', 'Google', 64000)")
    if err != nil{
        panic(err)
    }
    fmt.Println(result.LastInsertId())  // id добавленного объекта
    fmt.Println(result.RowsAffected())  // количество затронутых строк
}