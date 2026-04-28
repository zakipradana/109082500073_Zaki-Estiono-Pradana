package main
import "fmt"

func main(){
	var ikan [1000] float64
	var total, jumlah float64
	var x, y , wadah int

	fmt.Scan(&x, &y)

	for i:=0; i<x; i++ {
		fmt.Scan(&ikan[i])
	}
	fmt.Println()

	wadah = 1
	for i:=0; i<x; i+=y {
		total = 0
		jumlah = 0
		
		for j:=0; j<y; j++ {
			if i + j >= x {
				break
			}else{
				total += ikan[i+j]
				jumlah++
			}
		}
		fmt.Printf("Total Berat Ikan  diWadah %d: %.2f\n", wadah, total)
		wadah++
	}

	
	for i := 0; i < x; i+=y {
		total = 0
		jumlah = 0
		for j := 0; j < y; j++ {
			if i + j >= x {
				break
			} else {
				total += ikan[i+j]
				jumlah++
			}
			
		}
		fmt.Printf("rata-rata tiap wadah: %.2f\n", total/jumlah)
       	wadah++
	}
}