# <h1 align="center">Laporan Praktikum Modul 4 - Algoritma Pemograman 2 </h1>
<p align="center">Zaki Estiono Pradana - 109082500073</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). 
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.
#### permutasikombinasi.go

```go
package main
import "fmt"

func faktor(n int, hasil *int){
    *hasil=1
    for i:= 1; i<=n; i++{
        *hasil *=i
    }
   
}
func permutasi(n, r int, hasil *int){
    var nfak, nrfak int
    faktor(n, &nfak)
    faktor(n - r, &nrfak)
    *hasil = nfak/nrfak
}
func kombinasi(n, r int, hasil *int){
    var nfak, nrfak, rfak int
    faktor(n, &nfak)
    faktor(r, &rfak)
    faktor(n - r, &nrfak)
    *hasil = nfak / (rfak*nrfak)
}
func main() {
    var a, b, c, d, a1, b1, a2, b2 int
    fmt.Scan(&a, &b, &c, &d)

    permutasi (a,c, &a1)
    kombinasi (a,c, &b1)
    permutasi (b,d, &a2)
    kombinasi (b,d, &b2)

    fmt.Println(a1, b1)
    fmt.Println(a2, b2)

  
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek4/output/permutasikombinasi.png)

##### Output 
[penjelasan]
Kode pemograman diatas merupakan kode sederhana untuk menghitung faktor, permutasi dan kombinasi dengan function prosedur. Terdapat tiga function yaitu faktor, permutasi dan kombinasi dengan paramater Pass by Reference (pointer) pada variabel hasil. Untuk menghitung faktor diperlukan perulangan dengan nilai inisiasi awal 1,bertambah setiap iterasinya 1 dan berhenti sampai nilai yang diinputkan(n). Kemudian disimpan pada variabel hasil sebagai pointer dan dilakukan perkalian dengan "i" Pada function permutasi terdapat variabel untuk menyimpan nilai n faktorial "nfak" dan nilai dari n-r faktorial "nrfak". Kemudian dilakukan penghitungan operasi bagi pada variabel hasil dengan paramater pointer dan disimpan pada variabel tersebut. function kombinasi juga sama hanya saja bertambah variabel "rfak" sebagai menyimpan nilai dari r faktorial. pada kombinasi juga dilakukan pengoperasian kombinasi pada variabel hasil sebagai parameter pointer. Terdapat input yang dideklarasikan pada variabel "a, b, c, d" dengan tipe data integer, pemanggilan function dilakukan dua kali pada masing-masing function permutasi dan kombinasi dengan variabel "a1, a2" sebagai penyimpan hasil dari permutasi, "b1, b2" sebagai penyimpan hasil kombinasi. terdapat ketentuan permutasi dan kombinasi yaitu nilai pada variabel "a, c" & "b,d" maka output menghasilkan dua nilai permutasi "a1, b1" dan dua kombinasi "a2, b2".

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### gema.go

```go
package main
import "fmt"


func hitungSkor(soal, skor *int){

    *soal = 0
        *skor = 0
        var w int
        for i:=0; i <8; i++{
            fmt.Scan(&w)
            
            if w< 301{
                *soal +=1
                *skor +=w
            }
        }
    }
func main(){
    var nama, pemenang string
    var soal, skor, jumlahsoal, jumlahskor int
    jumlahskor = 9999

    fmt.Scan(&nama)
    for nama != "Selesai"{
        hitungSkor(&soal, &skor)

        if soal > jumlahsoal || (soal == jumlahsoal && skor < jumlahskor ){
            jumlahsoal = soal
            jumlahskor = skor
            pemenang = nama
        }
        fmt.Scan(&nama)
    }
    fmt.Println(pemenang, jumlahsoal, jumlahskor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek/output/gema.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana untuk menentukan pemenang dari sebuah perlombaan, Pemenang ditentukan dari banykanya menyelesaikan soal dengan waktu tercepat dengan jumlah soal 8 soal. dengan menggunakan prosedur untuk menghitung jumlah soal dan jumlah skor dari perolehan menjawab soal sebanyaknya dan tercepat. Dengan nama function "hitungSkor" dan parameter "soal, skor" menggunakan pointer bertipe data integer, Inisiasi awal ditunjukan pada nilai awal jumlah soal dan jumlah skor "0". Perluangan berjalan sebanyak 8 kali dengan nilai awal inisiasi 0 dan bertambah 1 setiap iterasinya, setiap jumlah iterasi disimpan pada variabel "w" bertipe data integer, dengan percabangan ketika nilai pada variabel "w" < 301 (variabel w diartikan sebagai waktu yang ditempuh untuk mengerjakan soal) maka jumlah soal bertambah 1 dan jumlah skor ditambah dengan nilai "w", namun jika terdapat nilai 301 maka nilai 301 tidak terhitung. Untuk menentukan pemenang dilakukan pada function main atau diluar prosedur namun tetap memanggil function "hitungskor" sebagai penentu skor tertinggi, terdapat variabel " nama, pemenang" dengan tipe data string, dan variabel "soal, skor, jumlahsoal, jumlahskor" dengan tipe data integer. nilai maksimal jumlahskor diset di 9999. Input menujukan nama peserta lomba, karena untuk mengahkiri dengan menginputkan selesai, maka menggunakan for untuk program ini terus berjalan hingga mnenginputkan "Selesai" Ketika menginputkan Selesai maka menampilkan pemenang. untuk menentukan pemenang dilakukan percabangan ketika soal > jumlahsoal atau soal = jumlahsoal dan skor < jumlahskor, artinya prioritas utamanya jumlahsoal yang terjawab, namun ketika seri dihitung dari jumlah skor(waktu tersingkat) dengan aksi Jumlahsoal diisi debngan banyaknya jumlah soal terbaru, jumlahskor juga diisi dengan banyaknya jumlahskor terbaru, jadi ketika terdapat 2 peserta point nilai akan selalu update pada aksi tersebut dimana sesuai dengan ketentuan pada percabangan. Untuk menampilkan pemenang maka memanggil variabel pemenang, dengan  menampilkan perolehan jumlah soal dan jumlah skor
