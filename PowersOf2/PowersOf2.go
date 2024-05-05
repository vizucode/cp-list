package powersof2

func PowersOf2(n int) (resp []uint) {

	// power := 1
	// resp = append(resp, 1)
	// for i := 1; i <= n; i++ {
	// 	power = power * 2
	// 	resp = append(resp, uint(power))
	// }

	for i := 0; i <= n; i++ {
		resp = append(resp, 1<<i)
	}

	return resp
}
