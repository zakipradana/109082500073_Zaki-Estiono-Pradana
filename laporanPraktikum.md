# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Zaki Estiono Pradana] - [109082500073]</p>

## Unguided 

### 1. [string.go]
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

