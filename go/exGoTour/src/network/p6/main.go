package main
import (
    "fmt"
    "net/http"
    "io"
    "os"
    // "time"
)
func main() { 
    client := &http.Client{
        // Timeout: 5 * time.Second,
    } 
    resp, err := client.Get("https://telecom.tm") 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close() 
    io.Copy(os.Stdout, resp.Body)
}