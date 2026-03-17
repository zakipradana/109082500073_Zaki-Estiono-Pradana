package main
import "fmt"

func faktor(x int)int{
    hasil:=1
    for i:= 1; i<=x;i++{
        hasil=hasil * i
    }
    return hasil
}
func permutasi(x, r int)int{
    return faktor(x)/faktor(x-r)
}
func kombinasi(x, r int)int{
    return faktor(x)/(faktor(r)*faktor(x-r))
}
func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(permutasi(a, c), kombinasi(a, c))
    fmt.Println(permutasi(b, d), kombinasi(b, d))

  
}