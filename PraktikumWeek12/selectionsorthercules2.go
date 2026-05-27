	package main
	import "fmt"

	type arr [1000]int

	func assecending(a *arr, n int){
		var t, i, j, idx_min int
		i = 1
		for i <= n-1 {
			idx_min = i - 1
			j = i
			for j < n {
				if a[idx_min] > a[j] {
					idx_min = j
				}
				j++
			}
			t = a[idx_min]
			a[idx_min] = a[i-1]
			a[i-1] = t
			i++
		}
	}

	func desecending(a *arr, n int){
		var t, i, j, idx_max int
		i = 1
		for i <= n-1 {
			idx_max = i - 1
			j = i
			for j < n {
				if a[idx_max] < a[j] {
					idx_max = j
				}
				j++
			}
			t = a[idx_max]
			a[idx_max] = a[i-1]
			a[i-1] = t
			i++
		}
	}

	func main(){
		var i,j, k, m, n int
		var ganjil, genap arr
		var jumlah_ganjil, jumlah_genap int

		
		fmt.Print("Masukan banyaknya daerah: ")
		fmt.Scan(&n)
		i = 0
		for i < n {
			fmt.Print("Input no daerah ke-", i+1, ": ")
			fmt.Scan(&m)
			jumlah_ganjil = 0
			jumlah_genap = 0

			j = 0

			for j < m {
				var nomor int
				fmt.Scan(&nomor)
				if nomor % 2 != 0 {
					ganjil[jumlah_ganjil] = nomor
					jumlah_ganjil++
				} else {
					genap[jumlah_genap] = nomor
					jumlah_genap++
				}
				j++
			}
			assecending(&ganjil, jumlah_ganjil)
			desecending(&genap, jumlah_genap)

			fmt.Print("Output: ")
			k = 0
			for k < jumlah_ganjil+jumlah_genap {
				if k > 0 {
					fmt.Print(" ")
				}
				if k < jumlah_ganjil {
					fmt.Print(ganjil[k])
				} else {
					fmt.Print(genap[k-jumlah_ganjil])
				}
				k++
			}
			fmt.Println()
			i++
		}
	}