package main

import "fmt"

type arr [1000]int

type data struct {
	n int
	a arr
}

func insertionsort(a *data) {
	var i, j, sementara int

	i = 1
	for i < a.n {
		sementara = a.a[i]
		j = i - 1

		for j >= 0 && a.a[j] > sementara {
			a.a[j+1] = a.a[j]
			j--
		}

		a.a[j+1] = sementara
		i++
	}
}

func cekjarak(a *data) {
	var i, jarak int
	var status bool

	if a.n <= 1 {
		fmt.Println("Data Berjarak 0")
		return
	}

	jarak = a.a[1] - a.a[0]
	status = true

	for i = 2; i < a.n; i++ {
		if a.a[i]-a.a[i-1] != jarak {
			status = false
		}
	}

	if status {
		fmt.Printf("Data Berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func main() {
	var a data
	var x int

	a.n = 0

	fmt.Scan(&x)

	for x >= 0 {
		a.a[a.n] = x
		a.n++

		fmt.Scan(&x)
	}

	insertionsort(&a)

	for i := 0; i < a.n; i++ {
		fmt.Print(a.a[i], " ")
	}
	fmt.Println()

	cekjarak(&a)
}