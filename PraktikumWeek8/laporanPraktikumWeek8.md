# <h1 align="center">Laporan Praktikum Modul 10  </h1>
<p align="center">[Zaki Estiono Pradana] - [109082500073]</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

#### kelinci.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek8/output/kelinci.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menentukan berat terkecil dan terbesar dari sekumpulan kelinci. Pada program ini terdapat tipe data bentukan dengan nama data kelinci. Didalam type data tersebut terdapat variabel berat dengan array berindeks 1000 bertipe data float64, dan variabel jumlah dengan type data integer. Kemudian terdapat dalam fungsi utama terdapat variabel data dengan tipe data kelinci sebagai semua objek dalam struct. terdapat input untuk menyatakan jumlah kelinci yang dideklarsikan ke variabel jumlah. Perulangan pertama berfungsi untuk menginput berat kelinci disetiap kelinci dan berhenti sampai jumlah kelinci yang di inputkan. Variabel min dan max sebagai variabel penampung nilai min dan maximal dengan nilai awal 0. Untuk menentukan terbesar dan terkecil digunakan perulangan dengan membandingkan menggunakan percabangan. data berat akan di bandingkan dengan setiap iterasi sebagai contoh 5 3 8 1 4 maka nilai min dan max awal menjadi 5 kemudian menghitung iterasi selanjutnya yaitu 3, lalu apakah 3 < 5 maka menjadi nilai min terbaru, selanjutnya nilai max awalnya 5 akan dibandingkan dengan 8, karena 8 memenuhi ketentuan lebih besar maka 8 sebagai nilai max, terus berulang hingga iterasi tersebut selesai sesuai dengan jumlah banyaknya kelinci. Output akan menyatakan nilai min dan nilai max.

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2).Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### arraybilanganbulat.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek8/output/tarifikan.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi mengghitung total berat ikan disetiap wadah dan rata-rata berat ikan disetiap wadah. input x sebagai banyaknya ikan dan input y sebagai kapasitas ikan di wadah. terdapat perulangan pertama sebagai input berat masing-masing ikan. Perulangan kedua untuk menentukan total berat ikan ditiap wadah. dengan iterasi awal(j) 0, berhenti hingga sebelum y sebagai kapasitas wadah. kemudian terdapat percabangan untuk menentukan jumlah ikan dan tidak berlebih di masing-masing wadah, hal ini berfungsi ketika jumlah ikan ganjil dimasukan kedalam wadah dengan kapasitas genap atau yang menyebabkan satu wadah isinya tidak sama dengan wadah lainya. setelah itu menjumlahkan total berat ikan di setiap wadah dengan menggunakan menjumlahkan berat ikan setiap iterasinya dengan ketentuan jumlah kapasitas wadahnya. selanjutnya terdapat output yang menyatakan total berat ikan di setiap wadahnya. Untuk menentukan rata-rata berat ikan disetiap wadah cara kerjanya sama seperti menentukan berat ikan disetiap wadah hanya saja di rata-rata tinggal membagi berat ikan dengan jumlah ikan di setiap wadahnya.

### 3. SPos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

#### balita.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek8/output/balita.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menentukan berat maksimum, minimum, dan rata-rata berat balita. terdapat prosedur untuk menghitung maksimum dan minimum berat balita dengan parameter arrberat dengan type data Arrbalita berarray maximal di 100, n bertipe integer, bMin, bMax dengan tipe bertipe data float dan berpointer. nilai awal bMax dan bMin diset awal 0. kemudian perulangan berfngsi untuk menetukan berat maximum dan minimun dengan percabangan, Jika variabel arrBerat kurang dari min maka akan ditentukan sebagai nilai min, namun jika lebih berat maka ditentukan sebagai terberat, sebagai conoth jumlah balita 3 dan beratnya 1, 2, 3. maka akan menghitung dari indeks 0 karena 1 bukan < 1 maka akan nilai min di set 1 dan nilai max juga diset 1. mengecek indeks 1 karena 2 < 1 salah maka nilai min menjadi 1 dan nilai max menjadi 2 karena 2 > 1, selanuutnya mengecek indeks 3 karena bernilai 3 maka 3<1 salah maka nilai min tetap 1 dan nilai max menjadi 3 karena 3 > 2. kemudian fungsi rerata dengan parameter arrBerat bertipe ArrBalita dengan berpointer dan n bertipe data integer dan megembalikan nilai real atau float64, pada fungsi ini terdapat tambahan variabel lokal yaitu total dengan nilai awal 0. untuk menentukan total berat menggunakan perulangan dengan iterasi awal 0 dan berhenti hingga n sebagai banyaknya balita. kemudian disetiap iterasi / indeks akan dijumlahkan, indeks 1 akan dijumlahkan indeks 2 akan menghasilkan sebuah nilai dan menjumlahkan lagi dengan indeks / iterasi selanjutnya. kemudian mengembalikan nilainya atau return dengan membagikan nilai total dengan n sebagai banyaknya balita. Pada func main terdapat variabel balita dengan type data ArrBalita berarray max 100, n dengan type data integer sebagai jumlah banyaknya balita, rata, bMin, bMax bertipe data float64 sebagai penampung nilai. input jumlah banyaknya balita dengan deklarasi variabel n dan unutk menginputkan berat balita menggunakan perulangan dengan iterasi awal 0 berhenti hingga nilai n dan bertambah 1 setiap iterasinya yang dideklarasikan ke variabel balita yang ditampung pada array. kemudian memanggil fungsi hitungminmax dan fungsi rerata, dan output menyatakan nilai minimum maximum dan rerata.