package main
import "fmt"

type waktu struct{
	jam, menit, detik int
}
type DataRental struct{
	nama, jenisnaga string
	waktupinjam, waktukembali waktu
}


func main(){
	var data DataRental
	var tarif, total, durasi, jam, menit, detik int

	fmt.Scan(&data.nama)
	fmt.Scan(&data.jenisnaga)
	fmt.Scan(&data.waktupinjam.jam, &data.waktupinjam.menit, &data.waktupinjam.detik,)
	fmt.Scan(&data.waktukembali.jam, &data.waktukembali.menit, &data.waktukembali.detik)

	pinjam:= data.waktupinjam.jam*3600 + data.waktupinjam.menit*60+data.waktupinjam.detik
	kembali := data.waktukembali.jam*3600 + data.waktukembali.menit*60 + data.waktukembali.detik

	durasi = kembali - pinjam

	jam = durasi / 3600
	menit = durasi % 3600 / 60
	detik = durasi % 3600 % 60

	if data.jenisnaga == "Naga_Gembul"{
		tarif = 500
	}else if data.jenisnaga == "Naga_Cungkring" {
		tarif = 200
	}
	total = durasi * tarif
	
	if durasi >= 3600{
		total += 15000
	}
	fmt.Println("yth. ", data.nama, "Anda menyewa Naga", data.jenisnaga, "selama", jam, "jam" , menit, " menit",  detik, " detik")
	fmt.Println("total tagihan + uang seblak( jika ada): Rp: ", total )
}