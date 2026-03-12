# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Zaki Estiono Pradana] - [109082500073]</p>

## Unguided 

### 1. [Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?]
#### string.go

```go
package main
import "fmt"
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek1/output/stringimage.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi membaca tiga inputan bertipe data string dan menggeser satu langkah output tersebut sebagai output akhir. Variabel "satu, dua, tiga" merupakan variabel bertipe data string, variabel "temp" bertipe data string sebagai variabel temporary yang berfungsi untuk menyimpan sementera. Ketika menginputakan "aku", "dia", "kamu" maka output awal akan muncul "aku dia kamu" sesuai dengan urutan disimpan divariabel "satu, dua, tiga". Kemudian output akhir akan menggeser satu langkah, dimana variabel "temp" akan menyimpan dari data variabel "satu" kemudian berganti, variabel "satu" menyimpan dari data variabel "dua", variabel "dua" akan menyimpan variabel "tiga" kemudian variabel "tiga" akan menyimpan dari variabel "temp" yang isinya variabel "satu" sehingga output awal akan muncul "dia kamu aku"


### 2. [Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakann praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi  sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.]
#### warnakimia.go

```go
package main
import "fmt"

func main() {
    var w1, w2, w3, w4 string
    var i int
    var status bool
    
    status = true
    
    for i = 1; i <=5; i++{
        fmt.Printf("Percobaan ke %d: ", i)
        fmt.Scan(&w1, &w2, &w3, &w4)
        if !(w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu"){
            status = false
        }
    }
    fmt.Println("BERHASIL: ", status)
} 
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek1/output/warnakimia.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk mencatat hasil percobaan kimia sebanyak 5 kali percobaan dengan 4 tabung reaksi dengan ketentuan urutan warna merah, kuning, hijau, ungu. percobaan dinyatakan berhasil apabila susunan warna sebanyak 5 kali sama semua sesuai dengan ketentuan. Dengan menggunakan variabel "w1", "w2", "w3", "w4" dengan tipe data string, variabel i dengan tipe data integer, dan variabel "status' dengan tipe data boolean. Nilai kebenaran awal variabel status yaitu "true". Untuk melalukan inputan berulangan sebanyak 5 kali maka menggunakan perulangan, dimana inisiasi dimulai dari 1 dan berhenti ketika sebanyak 5 kali, dimana setiap inisiasi bertambah 1. aksi dari perulangan tersebut yaitu melakukan inputan percobaan yang disimpan ke variabel "w1", "w2", "w3", "w4", sampai dengan 5 kali perulangan. inputan bernilai true atau benar ketika sesuai dengan ketentuan warna yaitu merah, kuning, hijau, ungu. ketika salah satu percobaan terdapat input warna berbeda maka output berhasil akan bernilai false, namun jika sesuai maka bernilai true.

### 2. [PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.]
#### warnakimia.go

```go
    package main
import "fmt"

func main() {
    var Bparsel, beratkg, beratgr, biayagr, biayakg, Tbiaya int
    fmt.Print("Masukan Berat Parsel (gram): ")
    fmt.Scan(&Bparsel)
    
    beratkg = Bparsel / 1000
    beratgr = Bparsel % 1000
    fmt.Println("Detail berat: ", beratkg, "kg + ", beratgr, " gr")
    biayakg = beratkg * 10000
    
    if beratkg > 10{
        biayagr = 0
    }else if beratgr >= 500{
        biayagr = 5 * beratgr
    }else{
        biayagr = 15 * beratgr
    }
    fmt.Println("Detail Biaya: Rp. ", biayakg, " + Rp. ", biayagr )
    Tbiaya = biayakg + biayagr
    fmt.Println("Total biaya: ", Tbiaya)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/zakipradana/109082500073_Zaki-Estiono-Pradana/blob/main/PraktikumWeek1/output/tarifpos.png)

##### Output 
[penjelasan]
Kode pemograman ditas merupakan pemograman sederhana yang berfungsi untuk menentukan biaya tarif paket. Dengan beberapa ketentuan yaitu:
    1. Biaya jasa pengiriman adalah Rp. 10.000,- per kg.
    2. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram.
    3. Jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram.
    4. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
Pemograman diatas terdapat variabel “Bparsel, beratkg, beratgr, biayagr, biayakg, Tbiaya”, Dengan tipe data Integer, Variabel “Bparsel” untuk mendeklarasikan inputan nominal berat parsel dalam satuan gram, Variabel “beratkg” untuk menghitung konversi berat dari gram ke kilogram dengan rumus “Bparsel / 1000”, Variabel “beratgr” untuk menghitung sisa dari konversi gram ke kilogram dengan satuan gram dengran rumus “Bparsel % 1000”, Variabel “biayakg” digunakan untuk menghitung total biaya dalam satuan kg, kemudian terdapat 3 percabangan. Dengan ketentuan diatas maka pemograman menggunakan percabangan untuk menentukan kententuan berat. Percabangan pertama dengan kondisi beratkg lebih dari 10 kg maka biaya hanya di hitung biaya perkg, Percabangan 2 dengan kondisi jika variabel “beratgr” atau sisa berat atau berat lebih dalam satuan gram di sebuah parsel lebih dari sama dengan 500 gram maka terdapat aksi yaitu menghitung biaya lebih dengan rumus “ 5 * beratgr” yaitu dengan biaya 5 rupiah pergram, Percabangan 3 langsung menghitung dengan catatan semua sisa berat tidak masuk dalam syarat percabangan 1 dan 2 yang berarti kurang dari 500 gram maka terdapat aksi “15 * beratgr” yaitu dengan biaya 5 rupiah pergram, Kemudian output detail biaya menampilkan hasil dari penjumlahan biaya beratkg dan biaya berat lebih, setelah dijumlahkan maka terdapat output total biaya keseluruhan paket tersebut.