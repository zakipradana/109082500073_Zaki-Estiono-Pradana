package main
import "fmt"
import "math"

type titik struct{
    x, y float64
}

type lingkaran struct{
    pusat titik
    r float64
}z

func akar(x float64)float64{
	return math.Sqrt(x)
}

func jarak (p,q titik)float64{
	x:= (p.x-q.x) * (p.x-q.x)
	y:= (p.y-q.y) * (p.y-q.y)
	return akar(x+y)
}

func didalam(c lingkaran, p titik)bool{
	return jarak(c.pusat, p) <= c.r
}

func main (){
	var l1, l2 lingkaran
    var p titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	lingkaran1:= didalam(l1, p)
	lingkaran2:= didalam(l2, p)

	if lingkaran1 && lingkaran2{
		fmt.Println("Titik Berada di dalam lingkaran 1 dan 2")
	}else if lingkaran1{
		fmt.Println("Titik Berada di dalam lingkaran 1")

	}else if lingkaran2{
		fmt.Println("Titik Berada di dalam lingkaran 2")
	}else{
		fmt.Println("Titik Berada di luar lingkaran 1 dan 2")
	}
}