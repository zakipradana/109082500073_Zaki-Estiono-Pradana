package main
import "fmt"

type ArrBalita [100] float64

func hitungMinMax(arrBerat ArrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		} else if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat *ArrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(n)
}


func main(){
	var balita ArrBalita
	var n int
	var rata, bMin, bMax float64
	

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&balita[i])
	}

	hitungMinMax(balita, n, &bMin, &bMax)
	rata = rerata(&balita, n)
	fmt.Printf("Berat Balita Minimum: %.2fkg\n", bMin)
	fmt.Printf("Berat Balita Maksimum: %.2fkg\n", bMax)
	fmt.Printf("Rerata Berat Balita: %.2fkg\n", rata)
	
}