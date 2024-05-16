package magicchange

import "fmt"

/*

	N = 2
	P = [1,2,3]
	Q = [2,1,4]


	input

	// jumlah operasi pertukaran yang terjadi
	T = 1

	// mengambil x dan y pada tiap baris dan kemudian menukarnya
	P [2]
	Q [2]

	output
	P = [1, 1]
	Q = [2, 2]

*/

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func MagicChange() {
	var (
		N  int
		T  int
		ar [2][1001]int
	)

	fmt.Scan(&N)             // Baca jumlah elemen
	for i := 0; i < 2; i++ { // Baca elemen array
		for j := 0; j < N; j++ {
			fmt.Scan(&ar[i][j])
		}
	}

	fmt.Scan(&T) // Baca jumlah pertanyaan
	for i := 0; i < T; i++ {
		var (
			buff1 string
			buff2 string
			x     int
			y     int
			p     int
			q     int
		)

		fmt.Scan(&buff1, &x, &buff2, &y) // Baca pertanyaan

		p = int(buff1[0]) - 'A' // Konversi karakter ke indeks
		q = int(buff2[0]) - 'A'
		x--
		y--

		Swap(&ar[p][x], &ar[q][y]) // Tukar elemen
	}

	for i := 0; i < 2; i++ { // Cetak hasil akhir array
		for j := 0; j < N; j++ {
			fmt.Print(ar[i][j])
			if j+1 < N {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
