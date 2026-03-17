package main
import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)
    sisa:= y%x
    fmt.Print(sisa)
  
}