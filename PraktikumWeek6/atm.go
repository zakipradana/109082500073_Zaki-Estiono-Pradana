package main
import "fmt"

type rekening string
type pin string
type uang float64
func main() {
    var inputrekening rekening
    var inputpin pin
    var saldo uang

    var nilaisaldo uang= 500000
    fmt.Print("Masukan rekening: ")
    fmt.Scan(&inputrekening)
    fmt.Print("Masukan Pin: ")
    fmt.Scan(&inputpin)

    if inputrekening == "889900"{
        
        if  inputpin == "123456"{
            fmt.Print("Masukan Nominal Yang akan ditarik: ")
            fmt.Scan(&saldo)
            if saldo <= nilaisaldo{
                nilaisaldo -= saldo
                fmt.Print("Sisa Saldo: ", nilaisaldo)
            }else{
                fmt.Print("Maaf Saldo anda Habis, Silahkan Bekerja lebih keras!")
            }

        }else{
        fmt.Print("Maaf Pin salah")
        }
    }else{
    fmt.Print("No rekening Salah")
    }
  
}