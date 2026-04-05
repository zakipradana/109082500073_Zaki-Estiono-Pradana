# <h1 align="center">Laporan Praktikum Modul 4 - Algoritma Pemograman 2 </h1>
<p align="center">Zaki Estiono Pradana - 109082500073</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n!/(n−r)!, sedangkan C(n, r) = n!/r!(n−r)!
#### permutasikombinasi.go

```go
package main
import "fmt"

func faktor(x int)int{
    hasil:=1
    for i:= 1; i<=x;i++{
        hasil=hasil * i
    }
    return hasil
}
func permutasi(x, r int)int{
    return faktor(x)/faktor(x-r)
}
func kombinasi(x, r int)int{
    return faktor(x)/(faktor(r)*faktor(x-r))
}
func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(permutasi(a, c), kombinasi(a, c))
    fmt.Println(permutasi(b, d), kombinasi(b, d))

  
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek2/output/permutasikombinasi.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menghitung permutasi dan kombinasi. Terdapat tiga function yaitu faktor, permutasi, dan kombainasi. function faktor berguna untuk menghitung faktor dari bilangan yang diinputkan sebagai contoh input bilangan 5 10 3 10, maka function faktor menghitung menjadi 5!, 10!, 3!, 10!. untuk menghitung faktor digunakan perulangan dengan ketentuan mengulang secara terus menerus sebanyak bilangan yang dinputkan dengan hasil dikalikan hasil inisasi dengan nilai hasil awal yaitu 1. sebagai contoh 3! maka menghitung dengan cara 1*1=1, 1*2=2, 2*3=6, 6*4=24, maka hasil 4!= 24. Function permutasi berfungsi untuk menghitung permutasi dari bilangan yang dinputkan sebagai contoh permutasi 5 terhadap 3 maka: 5!/(5!-3!)= 60.function kombinasi berfungsi untuk menghitung kombinasi daru bilangan yang diinputkan, sebagai contoh kombinasi 5 terhadap 3 maka: 5!/(2!*3!)=10. input terdiri dari 4 bilangan yang dideklrasikan oleh 4 variabel yaitu "a, b, c, d" dengan tipe data integer dan output menyatakan 2 baris dengan format permutasi, kombinasi.

### 2. Diberikan tiga buah fungsi matematika yaitu f(x) = x*x , g(x) = x − 2 dan h(x) = x +1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!
#### fungsi.go

```go
package main
import "fmt"

func f(x int)int{
    return x*x
}
func g(x int)int{
    return x-2
}
func h(x int)int{
    return x+1
}

func main (){
    var q,w,e int
    fmt.Scan(&q,&w,&e)
    fmt.Println(f(g(h(q))))
    fmt.Println(g(h(f(w))))
    fmt.Println(h(f(g(e))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek2/output/fungsi.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menghitung sebuah fungsi. Terdapat 3 function, function f berfungsi menghitung fungsi f(x), function g untuk menghitung fungsi g(x), function h untuk menghitung fungsi g(h). Terdapat 3 variabel "q, w, e" dengan tipe data integer untuk mendeklarasikan input sebuah bilangan. variabel "q" untuk operasi pertama, "w" untuk operasi kedua, dan "e" untuk operasi ketiga, Output pertama menghasilkan operasi (fogoh)(q), baris kedua(gohof)(w), dan baris ketiga (hofog)(e).

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### warnakimia.go

```go
package main
import "fmt"
import "math"

func jarak(a, b, c, d float64)float64{
    hasil:=math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
    return hasil
}
func didalam(cx, cy, r, x, y float64)bool{
    if jarak(cx, cy, x, y) <= r{
		return true
    }
    return false
}
func main() {
    var cx1, cy1, r1, cx2, cy2, r2, x, y int
    fmt.Scanln(&cx1, &cy1, &r1)
    fmt.Scanln(&cx2, &cy2, &r2)
    fmt.Scanln(&x, &y)

    ling1:= didalam(float64(cx1), float64(cy1), float64(r1), float64(x), float64(y))
    ling2:= didalam(float64(cx2), float64(cy2), float64(r2), float64(x), float64(y))

    if ling1 && ling2{
        fmt.Println("Titik di dalam lingkaran 1 dan 2")

    }else if ling1{
        fmt.Println("Titik di dalam lingkaran 1")

    }else if ling2{
        fmt.Println("Titik di dalam lingkaran 2")
    }else{
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}   
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek2/output/lingkaran.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana untuk menentukan sebuah titik (x,y) di sebuah lingkaran. Terdapat dua function yaitu function jarak yang berfungsi untuk menentukan titik pusat x, y. sedangkan function didalam untuk menentukan titik tersebut berada di lingkaran dalam atau luar. terdapat variabel "cx1, cy1, r1, cx2, cy2, r2, x, y" dengan tipe data integer berfungsi sebagai nilai input. kemudian terdapat ketetuan input lingkaran 1 ""cx1, cy1, r1, x, y" dan lingkaran 2 " "cx2, cy2, r2, x, y". Sebagai contoh input baris pertama 1 1 5 yang artinya lingkaran 1 dengan titik pusat (1 1), dan radius 5. Baris kedua 8 8 4 yang artinya lingkaran 2 dengan titik pusat (8 8), dan radius 4, dan baeis ketiga 2 2 merupakan titik sembarang (x, y).kemudian dilakukan menghitung titik pusat dengan function jarak = akar(2-1)^2+(2-1)=1.41 ,kemudian hasil tersebut ditentukan nilainya pada function didalam, ketika 1, 1, 5, 2, 2 memiliki hasil 1.41 dengan ketentuan <= r(5) artinya bernilai true. kemudian mengecek pada lingkaran 2 dengan function jarak = (2-8)^2+(2-8)^2=8.49 , lalu masuk pada function didalam ketika 8, 8, 4, 2, 2 memiliki hasil jarak 8.49 sehingga melebihi r(4) maka bernilai false. Setelah semua selesai dilakukan operasi maka masuk dalam tahap percabangan, pada contoh ini hanya lingkaran 1yang bernilai true maka output menyatakan berada di lingkaran 1.