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