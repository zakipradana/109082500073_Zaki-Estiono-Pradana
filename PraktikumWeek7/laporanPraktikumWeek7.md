# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">[Zaki Estiono Pradana] - [109082500073]</p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### lingkaranarray.go

```go
package main
import "fmt"
import "math"

type titik struct{
    x, y float64
}

type lingkaran struct{
    pusat titik
    r float64
}

func akar(x float64)float64{
	return math.Sqrt(x)
}

func jarak (p,q titik)float64{
	x:= (p.x-q.x) * (p.x-q.x)
	y:= (p.y-q.y) * (p.y-q.y)
	return akar(x+y)
}

func didalam(c lingkaran, p titik)bool{
	return jarak(c.pusat, p) <= c.r
}

func main (){
	var l1, l2 lingkaran
    var p titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	lingkaran1:= didalam(l1, p)
	lingkaran2:= didalam(l2, p)

	if lingkaran1 && lingkaran2{
		fmt.Println("Titik Berada di dalam lingkaran 1 dan 2")
	}else if lingkaran1{
		fmt.Println("Titik Berada di dalam lingkaran 1")

	}else if lingkaran2{
		fmt.Println("Titik Berada di dalam lingkaran 2")
	}else{
		fmt.Println("Titik Berada di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek7/output/lingkaranarray.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi menentukan titik didalam lingkaran atau diluar lingkaran. dengan dua tipe data bentukan yaitu titik dengan nama data x, dan y sebagai variabel bertipe data integer, dan lingkaran dengan nama data pusat bertipe data titk dan r bertipe data float64 atau real. untuk menentukan sebuah titik disuatu lingkaran diperulkan akar. Untuk mengoperasikan akar diprogram ini menggunakan fungsi dengan bertipe data float64 sebagai parameter. selanjutnya menghitung akar dalam fungsi akar tersebut menggunakan library math dengan bantuan math.Sqrt. kemudian terdapat fungsi jarak untuk menentukan jarak antar kedua titik dengan paramater p dan q menujukan sebuah titik koordinat. untuk menentukannya menggunakan perhitungan selisih dimasing-masing kordinat x dan yy, lalu di kuadartkan. Selanjutnya dikembalikan dengan mengakarkan hasil selisih tersebut. Kemudian fungsi didalam sebagai penentu sebuah titik disebuah lingakaran, dengan parameter c yang merupakan sebuah lingkaran sebagai pusat dan radiusnya dan p sebuah titik dan mengembalikan nilai bernilai boolean. Fungsi ini hanya bekerja mengembalikan nilai ke fungsi jarak dengan menghitung jarak antara pusat lingkaran dan jari-jari dengan ketentuan jarak titik ke pusat lebih kecil atau sama dengan radius mmaka titik di dalam atau di tepi lingkaran namun jika lebih besar maka titik diluar. Pada func main terdapat variabel l1, l2 sebagai lingkaran 1 dan lingkaran 2, dan p sebagai titik. terdapat input titik pusat lingkaran dan jari-jari dimasing masing lingkaran 1 dan 2 dibedakan dengan baris1 dan baris 2, kemudian dibaris ketiga menginputkan titik koordinat. Selanjutnya dilakukan pengecekan posisi titik apakah berda didalam lingkaran atau tidak disetiap lingkarannya hal ini akan dikirim ke masing-masing fungsi untuk dilakukan penghitungan dimana titik koordinat dan titik pusat akan dikirm ke fungsi jarak dan akar. Kemudian dilakukan pengecekan dengan percabangan untuk menentukan titik berada dilingkaran mana saja baik dilingkaran 1, lingkaran 2 atau diluar.

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
    a. Menampilkan keseluruhan isi dari array.
    b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
    c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
    d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
    e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
    f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
    g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
    h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

#### arraybilanganbulat.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek7/output/arraybilanganbulat.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk untuk menampung sekumpulan bilangan bulat. dengan beberapa ketentuan seperti di soal. Fungsi pertama terdapat akar yaitu untuk menghitung akar yang digunakan pada standar deviasi. Fungsi semua berfungsi untuk menampilkan seluruh value array yang diinputkan dari total indeks yang diinput, diprogram ini maximal indeks arraynya yaitu 100. untuk menampilkan menggunakan perulangan dengan nilai i dimulai dari 0 sebagai indeks pertama sampai i< n yaitu indeks terakhir, dan ditampilkan value setiap iterasinya (indexnya). Fungsi selanjutnya yaitu menampilkan array di indeks ganjil, dengan menggunakan perulangan seperti menampilkan semua isi array hanya saja iterasi awalnya dimulai dari 1 dan isetiap iterasinya bertambah 2. Menampilkan hanya indeks ganjil. Fungsi menampilkan indeks genap sama seperti fungsi menampilkan indeks ganjil hanya saja iterasi awalnya dimulai dari 0 dan setiap iterasinya bertambah 2. Fungsi kelipatan yaitu menampilkan elemen array pada indeks yang merupakan kelipatan dari sbeuah bilangan x dengan parameter arr berindeks 100, n sebagai parameter jumlah indeks yang terpakai dan x merupakan nilai kelipatan bertipe integer, Dengan sayarat nilai kelipatan tidak boleh 0. Untuk mengecek kelipatan menggunakan perulangan dimulai dari inisiasi bernilai 0, kemudian untuk menentukan bilangan kelipatan maka menggunakan modulus dengan membagi setiap iterasinya. Fungsi hapus berfungsi untuk menghapus elemen array pada indeks tertentu dan menampilkan yang telah dihapus. Difungsi ini menggunakan  pointer diparameter array dan nilai n sebagai jumlah elemen, Paramater z sebagai indeks yang ingin dihapus. untuk menghapusnya menggunakan perulangan dimulai dari nilai yang ingin dihapus dan setiap elemen indeksnya berganti ke indeks yang lebih kecil. untuk menampilkan value array yang telah dihapus tetap menggunakan perulangan. fungsi ratarata berfungsi untuk mencari nilai rata-rata dari value sebuah array. terdapat variabel lokal yaitu sum bertipe data integer berfungsi menampung total semua elemen dengan nilai awal 0. untuk menjumlahkannya menggunakan perulangan yang setiap aksinya menambahkan elemen array di setiap indeks yang kemudian dibagi dengan jumlah indeks. Fungsi deviasi berfungsi untuk melihat seberapa jauh sebaran sebuah data, untuk menghitungnya pada fungsi ini menggunakan rata-rata dari value sebuah aray kemudian dihitung selisih masing-masing value terhadap selisih dan dikuadratkan. selanjutnya menghitung variannya dengan total perhitungan seleisih lalu dibagi jumlah indeks. untuk menghitung final standar deviasi digunakan akar, nilai total selisih tadi akan dikirim kefungsi akar dan menghasilkan nilai deviasi. Kemudian fungsi frekuensi sebagai menghitung berapa kali suatu nilai muncul dalam array. untuk menghitung membutuhkan variabel lokal, difungsi ini bernama freq bertipe data integer dengan nilai awal 0. perulangan berfungsi untuk mengecek semua value di setiap indeks. untuk menghasilkan nilai frekuensi yang dicari digunakan percabangan dengan syarat value yang dicari maka frekuensi akan menambah 1.Pada fungsi utama berfungsi memanggil dan memasukan nilai pada fungsi fungsi tertentu, terdapat variabel n, i, x, z, cari bertipe data ineteger. Variabel n sebagai jumlah indeks yang dimasukan. Variabel i sebagai inisiasi perulangan yang befungsi menampilkan indeks di setiap elemen.  variabel x sebagai nilai kelipatan, variabel z sebagai nilai yang akan dihapus dan variabel cari berfungsu untuk mencari frekuensi value di sebuah array.

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

#### pertandinganbola.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek7/output/pertandinganbola.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. terdapat prosedur untuk menentukan pemenang dengan parameter skorA, SkorB, index beripe data integer KlubA, KlubB bertipe data string dan array dengan maksimal indeks 100. Syarat pemenang yaitu jika skorA > skorB maka yang menang klubA, jika skorB > skorA maka yang menang klubB, namun jika seri maka menyatakan Draw. Pada fungsi utama terdapat variabel SkorA, SKorB bertipe data integer sebagai input skor masing-masing klub, KlubA, KlubB bertipe data string berfungsi untuk menginputkan nama klub. variabel pemenang sebagai array untuk menyimpan hasil/pemenang tiap pertandingan, nilai inisialisasi dimulai dari 0. Kemudian terdapat perulangan untuk menginputkan Skor masing-masing klub dan berhenti ketika skorA atau skorB bernilai < 0 atau bernilai -1, jika berhenti maka memanggil prosedur penentupemenang. kemudian ditampilkan semua hasil yang sudah disimpan dengan perulangan.

### 3. sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
#### palindrom.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWee7/output/palindrom.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom. array bernama NMAX bertipe data integer dengan nilai maximum 127. Kemudian terdapat tipe baru bernama tabel dengan tipe data rune. Fungsi array berfungsi untuk menginputkan karakter ke dalam sebuah array dengan inisialisasi awal 0, variabel x sebagai variabel lokal, terdapat perulangan untuk menetukan pemberhentian input dengan kondisi menginputkan "." atau sebuah value arraynya penuh. Untuk menyimpan value array memiliki ketentuan dimana spasi dan enter tidak disimpan yang tersimpan hanya karakter saja. Kemudian fungsi cetak array berfungsi untuk menampilkan isi array karakter. untuk menampilkannya setiap value array digunakan perulangan dengan nilai insialisasi 0 sampai dengan nilai indeksnya dan setiap itersinya bertambah 1. Fungsi balikanArray berfungsi untuk mereverse value array didalam fungsi ini terdapat variabel baru sebagai variabel lokal yang berfungsi menyimpan array baru, proses membalikan digunakan perulangan dengan inisialisasi awal 0 hingga jumlah indeks yang diinputkan dengan setiap iterasinya bertambah 1. Variabel B sebagai variabel baru bekerja disini sebagai menyimpan hasil array baru dengan mengambil value array lama. Fungsi palindrom berfungsi untuk mengecek apakah value array yang diinputkan palindrom atau bukan. dengan paramater t sebagai array asli (bukan array baru hasil reverse), n sebagai jumlah indeks array dan bernilai true or false. terdapat variabel x sebagai variabel lokal yang berfungsi untuk digunakan sebagai salinan array, kemudian perulangan berfungsi untuk menyalin array dari array lama (t) ke array baru (x), kemudian terdapat pemanggilan fungsi balikan array ini berfungsi untuk membalik array namun bukn array lama yang dibalik melainkan array baru yaitu (x), untuk menentukan palindrom menggunakan pembanding disetiap valuenya dengan menggunakan perulangan dengan aksi array lama (t)dan array baru (x) harus sama jika berbeda langsung bernilai false atau bukan palindrom. pada fungsi utama berfungsi untuk memanggil semua fungsi dengan menginputkn sebuah input yang dideklrasikan pada variabel tab dan m.