# <h1 align="center">Laporan Praktikum Modul 14  </h1>
<p align="center">[Zaki Estiono Pradana] - [109082500073]</p>

## Unguided 

### 1. HBuatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.

Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.
Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

#### selectionsorthercules1.go

```go
package main

import "fmt"

type arr [1000]int

type data struct {
	n int
	a arr
}

func insertionsort(a *data) {
	var i, j, sementara int

	i = 1
	for i < a.n {
		sementara = a.a[i]
		j = i - 1

		for j >= 0 && a.a[j] > sementara {
			a.a[j+1] = a.a[j]
			j--
		}

		a.a[j+1] = sementara
		i++
	}
}

func cekjarak(a *data) {
	var i, jarak int
	var status bool

	if a.n <= 1 {
		fmt.Println("Data Berjarak 0")
		return
	}

	jarak = a.a[1] - a.a[0]
	status = true

	for i = 2; i < a.n; i++ {
		if a.a[i]-a.a[i-1] != jarak {
			status = false
		}
	}

	if status {
		fmt.Printf("Data Berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func main() {
	var a data
	var x int

	a.n = 0

	fmt.Scan(&x)

	for x >= 0 {
		a.a[a.n] = x
		a.n++

		fmt.Scan(&x)
	}

	insertionsort(&a)

	for i := 0; i < a.n; i++ {
		fmt.Print(a.a[i], " ")
	}
	fmt.Println()

	cekjarak(&a)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek13/output/insertsort.png)


##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menentukan selisih anatar elemen dan mengurutkannya, terdapat type data struct dengan variabel n, a dan kapasitas array 1000, Prosedur insertsort berfungsi untuk mengurutkan bilangan dari yang terkecil ke besar dengan metode insersort. Prosedur cek jarak untuk mengecek selisih antara setiap elemen, pada func main hanya berfungsi untuk menginputkan bilangan dan memanggil setiap prosedur

### 2. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu
perpustakaan. Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.
#### selectionsorthercules2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek13/output/perpustakaan.png)
![Screenshot Output Unguided 1_2](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek13/output/perpustakaan2.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk manajemen buku dalam perpustakaan dengan kapasitas array 7919 dan terdapat database yang disimpan pada 'Buku' dengan type data struct yang terdiri  D, judul, penulis, penerbit, eksemplar, tahun, rating. Terdapat beberapa prosedur seperti input data buku untuk menginput data buku, Prosedur cetak Terfavorit untuk menampilkan buku favorit, prosedur urut buku berfungsi mengurutkan buku dari tertinggi ke terendah, kemudian prosedur cetak5Terbaru berfungsi untuk mencetak 5 buku teratas setelah diurutkan, dan func cari buku untuk mencari buku berdasarkan binary search.

