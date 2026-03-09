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