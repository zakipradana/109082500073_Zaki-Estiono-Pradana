package main
import "fmt"

type Kelinci struct {
	berat [1000] float64
	jumlah int
}

func main() {
	var data Kelinci

	fmt.Print("Masukan jumlah Kelinci: ")
	fmt.Scan(&data.jumlah)

	for i := 0; i < data.jumlah; i++ {
		fmt.Printf("Masukan berat Kelinci ke-%d: ", i+1)
		fmt.Scan(&data.berat[i])
	}
	min := data.berat[0]
	max := data.berat[0]

	for i := 0; i < data.jumlah; i++ {
		if data.berat[i] < min {
			min = data.berat[i]
		}else if data.berat[i] > max {
			max = data.berat[i]
		}
		
	}
	fmt.Printf("Berat Terkecil: %.2f\n", min)
	fmt.Printf("Berat Terbesar: %.2f\n", max)
}