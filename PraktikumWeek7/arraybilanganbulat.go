package main

import "fmt"

func akar(x float64) float64 {
	if x == 0 {
		return 0
	}
	y := x
	for i := 0; i < 10; i++ {
		y = y - (y*y-x)/(2*y)
	}
	return y
}

// a
func semua(arr [100]int, n int) {
	fmt.Println("Menampilkan Semua isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// b
func ganjil(arr [100]int, n int) {
	fmt.Println("Menampilkan isi array indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// c
func genap(arr [100]int, n int) {
	fmt.Println("Menampilkan isi array indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// d
func kelipatan(arr [100]int, n int, x int) {
	fmt.Println("Menampilkan isi array kelipatan", x, ":")
	if x == 0 {
		fmt.Println("x tidak boleh 0")
		return
	}
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

// e
func hapus(arr *[100]int, n *int, z int) {
	for i := z; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--

	fmt.Println("\nArray setelah dihapus:")
	for i := 0; i < *n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("\n")
}

// f
func ratarata(arr [100]int, n int) {
	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	rata := float64(sum) / float64(n)
	fmt.Println("Rata-rata =", rata)
}

// g
func deviasi(arr [100]int, n int) {
	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	rata := float64(sum) / float64(n)

	var total float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		total += selisih * selisih
	}

	varians := total / float64(n)
	std := akar(varians)

	fmt.Println("Standar deviasi =", std)
}

// h
func frekuensi(arr [100]int, n int, cari int) {
	var freq int
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi =", freq)
}

func main() {
	var arr [100]int
	var n, i, x, z, cari int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("Elemen ke -%d: ", i)
		fmt.Scan(&arr[i])
	}

	semua(arr, n)
	ganjil(arr, n)
	genap(arr, n)

	fmt.Println()
	fmt.Print("Masukan bilangan kelipatan: ")
	fmt.Scan(&x)
	kelipatan(arr, n, x)

	fmt.Println()
	fmt.Print("Masukan indeks yang dihapus: ")
	fmt.Scan(&z)

	hapus(&arr, &n, z)
	fmt.Println("Hasil Telah dihapus:")
	semua(arr, n)

	fmt.Println()
	ratarata(arr, n)
	deviasi(arr, n)

	fmt.Print("Masukan bilangan yang dicari: ")
	fmt.Scan(&cari)
	frekuensi(arr, n, cari)
}