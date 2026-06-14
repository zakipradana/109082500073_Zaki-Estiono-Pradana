package main
import "fmt"

type arr [1000]int

func insertionsort(a *arr, n int){
	var i, j, sementara int

	i = 1
	for i <= n-1 {
		j = i
		sementara = a[i]
		for j > 0 && sementara < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = sementara
		i++
	}
}

func median(a *arr, n int) int {
	if n % 2 != 0 {
		return a[n/2]
	} else {
		return (a[n/2-1] + a[n/2]) / 2
	}
}	

func main(){
	var nomor arr
	var x, y int

	x = 0
	fmt.Print("Masukan data : ")
	fmt.Scan(&y)
	for y != -5313 {
		if y == 0 {
			insertionsort(&nomor, x)
			fmt.Println(median(&nomor, x))
		}else {
			nomor[x] = y
			x++
		}
		
		fmt.Scan(&y)
	}	
}
