package rotatematrix

import "fmt"

/*

	there was matrix M x N
	then we should do rotate matrix 90 degree

	1 2 3
	4 5 6
	7 8 9

	7 4 1
	8 5 2
	9 6 3

*/

func RotateMatrix() {

	var N int
	var M int

	fmt.Scan(&N)
	fmt.Scan(&M)

	// Initializing the matrix with size N x M
	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, M)
	}

	// input matrix N x M
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	// rotate and print | M x N
	for i := 0; i < M; i++ {
		// rotate from behind
		for j := N - 1; j >= 0; j-- {
			fmt.Print(matrix[j][i], " ")
		}
		fmt.Println()
	}

}
