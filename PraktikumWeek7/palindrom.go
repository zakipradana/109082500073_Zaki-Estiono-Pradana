package main
import "fmt"

const NMAX int = 127
type tabel [NMAX]rune


func isiArray(t *tabel, n *int){
	var x rune
	*n = 0

	for {
		fmt.Scanf("%c", &x)

		if x == '.' || *n >= NMAX {
			break
		}

		if x != ' ' && x != '\n' {
			(*t)[*n] = x
			(*n)++
		}
	}
}

func cetakArray(t tabel, n int){
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t tabel, n int) tabel {
	var B tabel

	for i := 0; i < n; i++ {
		B[i] = t[n-1-i]
	}

	return B
}
func palindrom(t tabel, n int) bool{
	var x tabel

	for i := 0; i < n; i++ {
		x[i] = t[i]
	}

	balikanArray(x, n)

	for i := 0; i < n; i++ {
		if t[i] != x[i] {
			return false
		}
	}
	return true
}

func main(){
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	fmt.Println("Palindrom ?", palindrom(tab, m))

	hasil := balikanArray(tab, m)

	fmt.Print("Reverse Teks : ")
	cetakArray(hasil, m)
}