package main
import "fmt"

func f(x int)int{
    return x*x
}
func g(x int)int{
    return x-2
}
func h(x int)int{
    return x+1
}

func main (){
    var q,w,e int
    fmt.Scan(&q,&w,&e)
    fmt.Println(f(g(h(q))))
    fmt.Println(g(h(f(w))))
    fmt.Println(h(f(g(e))))
}