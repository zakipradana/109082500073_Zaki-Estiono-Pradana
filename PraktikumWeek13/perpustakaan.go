package main
import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	var i int

	fmt.Print("Jumlah Buku : ")
	fmt.Scan(n)

	fmt.Println()
	fmt.Println("Masukkan Data Buku")
	fmt.Println("--------------------------------")

	for i = 0; i < *n; i++ {
		fmt.Println("Data Buku", i+1)

		fmt.Print("ID        : ")
		fmt.Scan(&pustaka[i].id)

		fmt.Print("Judul     : ")
		fmt.Scan(&pustaka[i].judul)

		fmt.Print("Penulis   : ")
		fmt.Scan(&pustaka[i].penulis)

		fmt.Print("Penerbit  : ")
		fmt.Scan(&pustaka[i].penerbit)

		fmt.Print("Eksemplar : ")
		fmt.Scan(&pustaka[i].eksemplar)

		fmt.Print("Tahun     : ")
		fmt.Scan(&pustaka[i].tahun)

		fmt.Print("Rating    : ")
		fmt.Scan(&pustaka[i].rating)

		fmt.Println()
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idx int

	idx = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
	}

	fmt.Println("Buku Terfavorit")
	fmt.Println("Judul    :", pustaka[idx].judul)
	fmt.Println("Penulis  :", pustaka[idx].penulis)
	fmt.Println("Penerbit :", pustaka[idx].penerbit)
	fmt.Println("Tahun    :", pustaka[idx].tahun)
	fmt.Println()
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var sementara Buku
	for i = 1; i < n; i++ {
		sementara = pustaka[i]
		j = i
		for j > 0 && pustaka[j-1].rating < sementara.rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = sementara
	}
}
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int
	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	fmt.Println("5 Buku Rating Tertinggi")
	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int
	var ketemu bool
	kiri = 0
	kanan = n - 1
	ketemu = false
	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if ketemu {
		fmt.Println("Data Buku Ditemukan")
		fmt.Println("Judul     :", pustaka[tengah].judul)
		fmt.Println("Penulis   :", pustaka[tengah].penulis)
		fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
		fmt.Println("Tahun     :", pustaka[tengah].tahun)
		fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
		fmt.Println("Rating    :", pustaka[tengah].rating)
	} else {
		fmt.Println("Buku Tidak Ditemukan")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	DaftarkanBuku(&pustaka, &n)

	fmt.Print("Masukkan Rating Buku Yang Dicari: ")
	fmt.Scan(&ratingCari)

	CetakTerfavorit(pustaka, n)	

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	CariBuku(pustaka, n, ratingCari)
}