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