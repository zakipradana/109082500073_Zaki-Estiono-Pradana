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
![Screenshot Output Unguided 1_1](output/image.png)
[penjelasan]
blablabalbakidiahi
