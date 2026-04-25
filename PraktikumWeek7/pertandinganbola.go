package main
import "fmt"

func penentupemenang(skorA, skorB, index int, klubA, klubB string, hasil *[100]string){
	if skorA > skorB {
		hasil[index]=klubA
	} else if skorB > skorA {
		hasil[index]=klubB
	} else {
		hasil[index]="Draw"
	}
}

func main() {
	var skorA, skorB, i int
	var klubA, klubB string
	var pemenang [100]string

	i = 0	
	fmt.Print("klub A: ")
	fmt.Scanln(&klubA)

	fmt.Print("klub B: ")
	fmt.Scanln(&klubB)

	for i = 0; i < 100; i++ {
		fmt.Printf("Pertandingan %d : ", i+1)
   		fmt.Scan(&skorA, &skorB)

    	if skorA < 0 || skorB < 0 {
        break
		}
		penentupemenang(skorA, skorB, i, klubA, klubB, &pemenang)

    }
	for j:=0; j<i; j++ {
		fmt.Printf("hasil %d: %s\n", j+1, pemenang[j])
	}
	fmt.Println("Pertandingan selesai")

}