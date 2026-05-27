# <h1 align="center">Laporan Praktikum Modul 14  </h1>
<p align="center">[Zaki Estiono Pradana] - [109082500073]</p>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 

Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-masing daerah.

#### selectionsorthercules1.go

```go
package main

import "fmt"

type arr [1000]int

func selectionsorthercules(a *arr, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if a[idx_min] > a[j] {
				idx_min = j
			}
			j++
		}
		t = a[idx_min]
		a[idx_min] = a[i-1]
		a[i-1] = t
		i++
	}
}

func main() {
	var i, m, n int
	var nomor_rumah arr

	fmt.Print("Masukan banyaknya daerah: ")
	fmt.Scan(&n)

	i = 0
	for i < n {
		fmt.Print("Input no daerah ke-", i+1, ": ")
		fmt.Scan(&m)
		j := 0
		for j < m {
			fmt.Scan(&nomor_rumah[j])
			j++
		}
		selectionsorthercules(&nomor_rumah, m)

		fmt.Print("Output: ")
		h := 0
		for h < m {
			if h > 0 {
				fmt.Print(" ")
			}
			fmt.Print(nomor_rumah[h])
			h++
		}
		fmt.Println()
		i++
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek12/output/selectionsorthercules1.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menentukan bilangan terkecil ke yang terbesar dengan metode selection sort. Terdapat array dengan kapasitas 1000 dengan tipe data integer. terdapat prosedur untuk menentukan bilangan terkecil ke yang terbesar, dan terdapat variabel lokal didalamnya dengan kententuan variabel "t" sebagai swap, "i" sebagai iterasi saat ini, "j" sebagai pencarian nilai terkecil dan "idx_min" sebagai indeks terkecil. terdapat while do sebagai perulangan sebagai pengecekan nilai terkecil, jika terdapat nilai terkecil maka akan disimpan pada "idx_min" kemudian dilakukan swap untuk memindahkan posisi. pada fungsi utama terdapat input sebagai penentuan banyaknya daerah yang dinput. Kemudian perulangan pertamaa digunakan sebagai input banyaknya nomor rumah dan akan mengulang sebanyak daerah yang di inputkan. setelah di inputkan kemudian akan dieksekusi oleh prosedur sorting dengan pemanggilan function. Selanjutkan ditampilkan karena inputnya sebanyak daerah maka output dilakukan perulangan sesuai banyaknya daerah yang diinputkan.

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

Format Masukan masih persis sama seperti sebelumnya.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
#### selectionsorthercules2.go

```go
	package main
	import "fmt"

	type arr [1000]int

	func assecending(a *arr, n int){
		var t, i, j, idx_min int
		i = 1
		for i <= n-1 {
			idx_min = i - 1
			j = i
			for j < n {
				if a[idx_min] > a[j] {
					idx_min = j
				}
				j++
			}
			t = a[idx_min]
			a[idx_min] = a[i-1]
			a[i-1] = t
			i++
		}
	}

	func desecending(a *arr, n int){
		var t, i, j, idx_max int
		i = 1
		for i <= n-1 {
			idx_max = i - 1
			j = i
			for j < n {
				if a[idx_max] < a[j] {
					idx_max = j
				}
				j++
			}
			t = a[idx_max]
			a[idx_max] = a[i-1]
			a[i-1] = t
			i++
		}
	}

	func main(){
		var i,j, k, m, n int
		var ganjil, genap arr
		var jumlah_ganjil, jumlah_genap int

		
		fmt.Print("Masukan banyaknya daerah: ")
		fmt.Scan(&n)
		i = 0
		for i < n {
			fmt.Print("Input no daerah ke-", i+1, ": ")
			fmt.Scan(&m)
			jumlah_ganjil = 0
			jumlah_genap = 0

			j = 0

			for j < m {
				var nomor int
				fmt.Scan(&nomor)
				if nomor % 2 != 0 {
					ganjil[jumlah_ganjil] = nomor
					jumlah_ganjil++
				} else {
					genap[jumlah_genap] = nomor
					jumlah_genap++
				}
				j++
			}
			assecending(&ganjil, jumlah_ganjil)
			desecending(&genap, jumlah_genap)

			fmt.Print("Output: ")
			k = 0
			for k < jumlah_ganjil+jumlah_genap {
				if k > 0 {
					fmt.Print(" ")
				}
				if k < jumlah_ganjil {
					fmt.Print(ganjil[k])
				} else {
					fmt.Print(genap[k-jumlah_ganjil])
				}
				k++
			}
			fmt.Println()
			i++
		}
	}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek12/output/selectionsorthercules2.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk mengurutkan nomor dari yang terkecil ke yang terbesar  dan yang terbesar ke terkecil dengan ketentuan mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Program ini menggunakan array dengan kapasitas 1000 bertipe data integer. Prosedur ascending berfungsi untuk mengurutkan fungsi terkecil ke terbesar menggunakan selection sort dengan menggunakan parameter pointer. Variabel lokal "t, i, j, idx_min" dengan type data integer berfungsi untuk menyimpan nilai pada function. dengan kententuan variabel "t" sebagai swap, "i" sebagai iterasi saat ini, "j" sebagai pencarian nilai terkecil dan "idx_min" sebagai indeks terkecil. erdapat while do sebagai perulangan sebagai pengecekan nilai terkecil, jika terdapat nilai terkecil maka akan disimpan pada "idx_min" kemudian dilakukan swap untuk memindahkan posisi. pada fungsi utama terdapat input sebagai penentuan banyaknya daerah yang dinput. Kemudian perulangan pertamaa digunakan sebagai input banyaknya nomor rumah dan akan mengulang sebanyak daerah yang di inputkan. Kemudian untuk menetukan nilai terbesar ke terkecil menggunakan descending, pada program ini menggunakan prosedur bernama desending, pada descending hanya berbeda pada penentuan nilai sebelum diswap dengan mengkategorikan kurang dari. Pada function main terdapat variabel i,j, k, m, n bertipe data integer dengan ketentuan "n" sebagai jumlah daerah. "m" sebagai nomor rumah, "i" sebagai iterasi perulangan daerah, "j" sebagai iterasi perulangan input, "k" sebagai iterasi perulangan output. variabel ganjil, genap dengan tipe data array merujuk ke array berkapasitas 1000. variabel jumlah_ganjil, jumlah_genap sebagai penyimpan nomor rumah sementara. perulangan pertama sebagai input nomor rumah di setiap daerah. untuk membedakan nomor genap dan ganjil menggunakan percabangan yang masing-masing ketentuan disimpan pada variabel jumlah_ganjil dan jumlah_genap. kemudian dilakukan sorting atau pengurutan sesuai ketentuan soal dimana asscending hanya ganjil dan descending hanya genap. untuk menampilkan hasilnya menggunakan perulangan dengan penggabungan ganjil dahulu baru genap.

### 3. SKompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

Keluaran adalah median yang diminta, satu data per baris.

#### medianinsertsort.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek12/output/medianinsertsort.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menentukan nilai median dan berhenti ketika menginputkan nilai -5313. Teradapat array dengan kapasitas 1000. Terdapat prosedur insertionsort berfungsi untuk mengurutkan nilai dari terkecil ke yang terbesar dengan metode insert. Pada function ini terdapat variabel lokal "i, j, sementara" dengan tipe data integer. variabel "i" sebagai inisiasi perulangan dengan nilai awal 1, variabel "j" sebagai inisiasi perulangan dengan nilai awal i, variabel "sementara" berfungsi untuk menyimpan nilai sementara yang telah di urutkan atau dipilih. padalangkah insert dengan menggunakann ketentuan dimana jika ketentuan tersebut dipenuhi maka dilakukan swap tanpa pengecekan angka didepannya. function median digunakan untuk mencari nilai tengah. Pada func main terdapat ketentuan dimana jika kita input -5313 maka inputan dinyatakan berakhir atau berhenti. untuk menampilkan dilakukan penguruttan terlebih dahulu baru dilakukan pencarian nilai tengah.