package main

import "fmt"

func main() {
	fmt.Println("================================")
	fmt.Println("         Soal 2. Faktor Bilangan")
	fmt.Println("               Disusun oleh:")
	fmt.Println("        Naufal Saifullah Yusuf")
	fmt.Println("================================")

	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Print("Faktor-faktor dari ", n, ": ")
	cetakFaktorRekursif(n, 1)
}

func cetakFaktorRekursif(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		cetakFaktorRekursif(n, i+1)
	}
}