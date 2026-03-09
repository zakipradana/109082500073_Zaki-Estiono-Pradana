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