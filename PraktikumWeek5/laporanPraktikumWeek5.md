# <h1 align="center">Laporan Praktikum Modul 5 - Algoritma Pemograman 2 </h1>
<p align="center">Zaki Estiono Pradana - 109082500073</p>

## Unguided 

### 1. MDeret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### fibonanci.go

```go
package main
import "fmt"

func fibonanci(n int)int{
    
     if n == 0 {
        return 0

     }
     if n == 1 {
        return 1
     }
     return fibonanci (n-1)+ fibonanci(n-2)
    
}

func main(){
    var n int
    fmt.Print("Masukan Bilangan: ")
    fmt.Scan (&n)

    fmt.Println(fibonanci(n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek5/output/fibonanci.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menghitung sebuah deret dengan nilai suku ke -0 adalah 0 dan nilai suku ke -1 adalah 1. dan untuk menghitung suku selanjutnya dengan rumus Sn = Sn−1 + Sn−2. Untuk menghitung deret tersebut menggunakan fungsi rekrusif. Pada pemograman ini fungsi rekrusif bernama "fibonanci" dengan "n" sebagao paramater bertype data integer. Karena nilai suku ke-0 adalah 0 dan suku ke-1 adalah 1 maka terdapat percabangan ketika "n" bernilai 0 maka output langsung bernilai 0 karena nilai pada percabangan tersebut dikembalikan ke 0, dan ketika "n" bernilai 1 maka output langsung bernilai 1 karena nilai percabangan tersebut dikembalikan ke 1. namun ketika input nilai suku bernilai lebih dari 1 maka langsung kembali kebagian rumus fibonanci fibonanci (n-1)+ fibonanci(n-2). kemudian pada fungsi utama terdapat variabel "n" bertipe data integer yang berfungsi untuk mendeklarasikan nilai input sebagai niali suku ke-n. Kemudian untuk menampilkan hasil deret fibonanci terdapat pemanggilan fungsi reksrusif "fibonanci" dengan argumen "n" pada output.


### 2. 2) Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

#### cetakbintang.go

```go
package main
import "fmt"

func cetakbintang(n int){
    if n == 0{
        return 
    }
    cetakbintang( n-1)
    for i:= 1;i<=n;i++{
        fmt.Print("*")

    }
    fmt.Println()
}
func main() {
	var n int
	fmt.Scan(&n)

	cetakbintang(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek5/output/cetakbintang.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menampilkan pola bintang dengan fungsi rekrusif. Fungsi rekrusif pada kode ini bernama cetak bintang dengan "n" sebagai paramater bertipe data integer. untuk menentukan pola dan menampilkan bintang seperti segitiga menggunakan percabangan dan perulangan, percabangan dilakukan untuk tidak mengeksekusi nilai "n" ketika bernilai 0, Ketika nilai "n" lebih dari 1 maka fungsi rekrusif berjalan dengan n-1, hal ini agar mencetak bintang dari yang terkecil hingga terbesar. Untuk  membentuk polanya digunakan perulangan dengan inisiasi "i" bernilai 1, dengan kondisi selama i <- n maka setiap iterasi akan bertambah 1, Setiap perulangan berjalan, Program akan mencetak simbol "*" sebanuyak nilai "i" dengan pola mengurut kebawah dan mengulang sampai perulangan selesai. Untuk menampilkan polanya maka perlu pemanggilan fungsi rekrusif pada fungsi utama main. Pada fungsi utama terdapat variabel "n" bertipe data integer yang berfungsi untuk mendeklrasikan nilai input. Kemudian terdapat pemanggilan fungsi rekrusif yaitu "cetakbintang", nilai "n" akan bersifat argumen dan diterukan ke fungsi rekrusif untuk dijalankan pada fungsi rekrusif.

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### faktor.go

```go
package main
import "fmt"

func faktor(n, i int) {
		if  i > n{
			return
		} 
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(n, i+1)
}

func main (){
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek5/output/faktor.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana untuk menghitung sebuah faktor dari sebuah bilangan dengan menggunakan fungsi rekrusif. Fungsi rekrusif pada kode ini diberi nama faktor dengan paramater "n" dan pembanding "i" dengan type data integer. Pada fungsi rekrusif ini terdapat percabangan sebagai kondisi berhenti. Kemudian percabangan selanjutnya untuk menghtiung faktor dengan rumus nilai "n" dimodulus "i" hasilnya 0 maka i termasuk faktor, sebagai contoh 6%1=0 maka 1 merupakan faktor, kemudian dilakukan pemanggilan rekrusif kembali untuk pengecekan ulang dan menampilkan nilai faktornya dengan menjumlahkan nilai i dengan 1. Kemudian pada func main berfungsi untuk mendeklarasikan input dengan variabel "n" tipe data integer. dan memanggil fungsi rekrusif dengan nilai inisiasi awal "1".