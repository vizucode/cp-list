package closefriendfactor

import (
	"fmt"
	"math"
)

/*

	define:
		factor person notation is X and Y
		person (i): Xi and Yi
		close friends factor (Ti,j) => |Xj - Xi| ^ D + |Yj - Yi| ^ D

	question
		maximum factor close friends
		minimum factor close friends

	input
		[3] => N (total of person)
		[1] => D (constant close friends factor)
		[0,1] => person i factor
		[1,1] => person j factor
		[10, 10] => person x factor
	output
		1 => min close friends factor
		19 => max close friends factors


	Test Case
	input
		N => 2
		D => 1
		Person 1 => 2 3
		Person 2 => 3 2

	Ti,j => 2

	Output
		0
		2

*/

func Absolute(x int) (resp int) {
	if x < 0 {
		resp = x * -1
		return resp
	}
	return x
}

func Pow(x, y int) (resp int) {
	resp = 1
	for i := 0; i < y; i++ {
		resp *= x
	}

	return resp
}

func CalcDistance(x1, x2, y1, y2, d int) (resp int) {
	T := Pow(Absolute(x2-x1), d) + Pow(Absolute(y2-y1), d)
	return T
}

func CloseFriendFactor() {
	var (
		N      int
		D      int
		Person [][]int
		x      int
		y      int
		max    int = math.MinInt
		min    int = math.MaxInt
	)

	fmt.Scan(&N)
	fmt.Scan(&D)

	for i := 0; i < N; i++ {
		fmt.Scan(&x, &y)
		Person = append(Person, []int{x, y})
	}

	for i := 0; i < len(Person); i++ {
		for j := i + 1; j < len(Person); j++ {
			distance := CalcDistance(Person[i][0], Person[j][0], Person[i][1], Person[j][1], D)
			if distance > max {
				max = distance
			}

			if distance < min {
				min = distance
			}
		}
	}

	fmt.Print(min, " ", max)
}
