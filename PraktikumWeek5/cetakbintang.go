package main
import "fmt"

func cetakbintang(n int){
    if n == 0{
        return 
    }
    cetakbintang( n-1)
    for i:= 1;i<=n;i++{
        fmt.Print("*")

    }
    fmt.Println()
}
func main() {
	var n int
	fmt.Scan(&n)

	cetakbintang(n)
}