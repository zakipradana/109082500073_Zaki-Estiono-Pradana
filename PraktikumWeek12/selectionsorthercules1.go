package main

import "fmt"

type arr [1000]int

func selectionsorthercules(a *arr, n int) {
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

func main() {
	var i, m, n int
	var nomor_rumah arr

	fmt.Print("Masukan banyaknya daerah: ")
	fmt.Scan(&n)

	i = 0
	for i < n {
		fmt.Print("Input no daerah ke-", i+1, ": ")
		fmt.Scan(&m)
		j := 0
		for j < m {
			fmt.Scan(&nomor_rumah[j])
			j++
		}
		selectionsorthercules(&nomor_rumah, m)

		fmt.Print("Output: ")
		h := 0
		for h < m {
			if h > 0 {
				fmt.Print(" ")
			}
			fmt.Print(nomor_rumah[h])
			h++
		}
		fmt.Println()
		i++
	}
}