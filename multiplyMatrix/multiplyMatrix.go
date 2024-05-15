package multiplymatrix

import (
	"fmt"
)

func MultiplyMatrix() {

	var N int
	var M int
	var P int

	fmt.Scan(&N)
	fmt.Scan(&M)
	fmt.Scan(&P)

	var matrix_a = make([][]int, N)
	for i := range matrix_a {
		matrix_a[i] = make([]int, M)
	}

	var matrix_b = make([][]int, M)
	for i := range matrix_b {
		matrix_b[i] = make([]int, P)
	}

	var matrix_c = make([][]int, len(matrix_a))
	for i := range matrix_c {
		matrix_c[i] = make([]int, len(matrix_b[0]))
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scan(&matrix_a[i][j])
		}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < P; j++ {
			fmt.Scan(&matrix_b[i][j])
		}
	}

	// baris a
	for i := 0; i < len(matrix_a); i++ {
		// kolom b
		for j := 0; j < len(matrix_b[0]); j++ {
			totalMatrix := 0
			for k := 0; k < len(matrix_b); k++ {
				baris_matrix_a := matrix_a[i][k]
				kolom_matrix_b := matrix_b[k][j]

				totalMatrix += baris_matrix_a * kolom_matrix_b
			}
			matrix_c[i][j] = totalMatrix
		}
	}

	for i := 0; i < len(matrix_c); i++ {
		for j := 0; j < len(matrix_c[0]); j++ {
			fmt.Print(matrix_c[i][j], " ")
		}
		fmt.Println()
	}

}
