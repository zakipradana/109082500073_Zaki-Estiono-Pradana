package main
import "fmt"
import "math"

func jarak(a, b, c, d float64)float64{
    hasil:=math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
    return hasil
}
func didalam(cx, cy, r, x, y float64)bool{
    if jarak(cx, cy, x, y) <= r{
		return true
    }
    return false
}
func main() {
    var cx1, cy1, r1, cx2, cy2, r2, x, y int
    fmt.Scanln(&cx1, &cy1, &r1)
    fmt.Scanln(&cx2, &cy2, &r2)
    fmt.Scanln(&x, &y)

    ling1:= didalam(float64(cx1), float64(cy1), float64(r1), float64(x), float64(y))
    ling2:= didalam(float64(cx2), float64(cy2), float64(r2), float64(x), float64(y))

    if ling1 && ling2{
        fmt.Println("Titik di dalam lingkaran 1 dan 2")

    }else if ling1{
        fmt.Println("Titik di dalam lingkaran 1")

    }else if ling2{
        fmt.Println("Titik di dalam lingkaran 2")
    }else{
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}   