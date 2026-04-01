package main
import "fmt"


func fibonanci(n int)int{
    
     if n == 0 {
        return 0

     }
     if n == 1 {
        return 1
     }
     return fibonanci (n-1)+ fibonanci(n-2)
    
}
func main(){
    var n int
    fmt.Print("Masukan Bilangan: ")
    fmt.Scan (&n)

    fmt.Println(fibonanci(n))
    
}
